const simple_binary = '1d081d1206666f6f6261721a04deadbeef220b088092b8c398feffffff01;'
const simple_json = {
  "Number": "29",
  "String": "foobar",
  "Bytes": "3q2+7w==",
  "Time": "0001-01-01T00:00:00Z"
}

const simple2_binary = '1c08afffffffffffffffff011a020001220b088092b8c398feffffff01;'
const simple2_json = {
  "Number": "-81",
  "String": "",
  "Bytes": "AAE=",
  "Time": "0001-01-01T00:00:00Z"
}

const big_binary = '5f0a1d081d1206666f6f6261721a04deadbeef220b088092b8c398feffffff010a1c08afffffffffffffffff011a020001220b088092b8c398feffffff011087ad4b1a1c6c6f6e67657220616e64206c6f6e67657220616e64206c6f6e676572;'
const big_json = {
  "Sub": [
    {
      "Number": "29",
      "String": "foobar",
      "Bytes": "3q2+7w==",
      "Time": "0001-01-01T00:00:00Z"
    },
    {
      "Number": "-81",
      "String": "",
      "Bytes": "AAE=",
      "Time": "0001-01-01T00:00:00Z"
    }
  ],
  "Count": "1234567",
  "Info": "longer and longer and longer"
}

