const { DynamoDBClient } = require("@aws-sdk/client-dynamodb");
const { DynamoDBDocumentClient, GetCommand, PutCommand } = require("@aws-sdk/lib-dynamodb");
const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');
const crypto = require('crypto')
const winston = require('winston');

const logger = winston.createLogger({
  // Log only if level is less than (meaning more severe) or equal to this
  level: "info",
  // format: winston.format.json(),
  defaultMeta: { service: 'user-service' },
  // Use timestamp and printf to create a standard log format
  format: winston.format.combine(
    winston.format.timestamp(),
    winston.format.json()
    // winston.format.printf(
    //   (info) => `${info.timestamp} ${info.level}: ${info.message}`
    // )
  ),
  // Log to the console and a file
  transports: [
    new winston.transports.Console(),
    new winston.transports.File({ filename: "logs/app.log" }),
  ],
});

const client = new DynamoDBClient({});
const docClient = DynamoDBDocumentClient.from(client);
const TABLE_NAME = "HypergridStatesDevNet";
const HSSN_RPC_DEFAULT = "https://exapi.devnet.hssn.sonic.game";

const PORT = process.env.PORT || 3000;
const app = express();
app.use(function(req, res, next) { 
  res.setHeader("Access-Control-Allow-Methods", "POST, PUT, OPTIONS, DELETE, GET, fetch");
  res.header("Access-Control-Allow-Origin", "*");
  res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
  res.header("Access-Control-Allow-Credentials", true);
  next();
});
app.use(bodyParser.json());
app.use(cors());

async function get_solana_account(rpc, address) {
  logger.info("get_solana_account", rpc, address);
  const url = rpc;
  const options = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      jsonrpc: "2.0",
      id: 1,
      method: "getAccountInfo",
      params: [
        address,
        {
          "encoding": "base64"
        }
      ]
    }),
  };

  try {
    let response = await fetch(url, options);
    let data = await response.text();
    logger.info(response.status, data);
    return [response.status, data];
  } catch (error) {
    logger.error(error);
    return [500, {error: error}];
  }
}

async function get_hypergrid_node(rpc, address) {
  logger.info("get_hypergrid_node", rpc, address);
  //get the nodes from cosmos via http request
  const url = rpc + "/hypergrid-ssn/hypergridssn/hypergrid_node/" + address;
  try {
    // let response = await axios.get(url);
    let response = await fetch(url);
    let data = await response.text();
    logger.info(response.status, data);
    return [response.status, data];
  } catch (error) {
    logger.error(error);
    return [500, {error: error}];
  }
}

async function add_to_db(rpc, address, version, status, data) {
  logger.info("add_to_db", rpc, address, version, status, data);
  //hash key = rpc + address + version
  hashkey = crypto.createHash('sha256').update(rpc + address + version).digest('hex');

  const params = {
    TableName: TABLE_NAME,
    Item: {
      hashkey: hashkey,
      rpc: rpc,
      address: address,
      version: version,
      status: status,
      data: data,
    },
  };

  try {
    const response  = await docClient.send(new PutCommand(params));
    logger.info(response);
  } catch (error) {
    logger.error(error);
  }
}

async function get_from_db(rpc, address, version) {
  logger.info("get_from_db", rpc, address, version);
  //hash key = rpc + address + version
  hashkey = crypto.createHash('sha256').update(rpc + address + version).digest('hex');
  const params = {
    TableName: TABLE_NAME,
    Key: {
      hashkey: hashkey,
    },
  };

  try {
    const response = await docClient.send(new GetCommand(params));
    logger.info(response);
    return response.Item;
  } catch (error) {
    logger.error(error);
    return null;
  }
}

app.post('/solana/GetAccountInfo', async (req, res) => {
  const {rpc, address, version} = req.body;
  if (!rpc || !address || !version) {
    res.status(400).send({state: 400, error:'invalid request'});
    return;
  }

  let resultdb = await get_from_db(rpc, address, version);
  if(resultdb != null) {
    let {status, data} = resultdb;
    if(!status) {
      status = 200;
    }
    res.status(status).type('application/json').send(data);
    return;
  }

  var result = await get_solana_account(rpc, address);
  if (result != null) {
    let [status, data] = result;
    await add_to_db(rpc, address, version, status, data);
    res.status(status).type('application/json').send(data);
  } else {
    res.status(404).send({state: 404, error: 'not found', result: null});
  }
});

app.post('/hssn/HypergridNode', async (req, res) => {
  var {rpc, address, version} = req.body;
  if(!rpc) {
    rpc = HSSN_RPC_DEFAULT;
  }
  if (!rpc || !address || !version) {
    res.status(400).send({state: 400, error:'invalid request'});
    return;
  }

  let resultdb = await get_from_db(rpc, address, version);
  if(resultdb != null) {
    let {status, data} = resultdb;
    if(!status) {
      status = 200;
    }
    res.status(status).type('application/json').send(data);
    return;
  }

  var result = await get_hypergrid_node(rpc, address);
  if (result != null) {
    let [status, data] = result;
    await add_to_db(rpc, address, version, status, data);
    res.status(status).type('application/json').send(data);
  } else {
    res.status(404).send({state: 404, error: 'not found', result: null});
  }
});

app.listen(PORT, () => {
  logger.info(`Server is running on port ${PORT}`);
});
