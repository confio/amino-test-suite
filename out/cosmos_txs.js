const unsigned_send_tx_binary = '65f0625dee0a482a2c87fa0a210a0c00cafe00deadbeef00cafe0012110a03424152120a39383736353433323130121f0a0a0123456789012345678912110a03424152120a3938373635343332313012150a0f0a03464f4f1208383736353433323110c0c407;'
const unsigned_send_tx_sign = '{"account_number":"1234","chain_id":"test-chain","fee":{"amount":[{"amount":"87654321","denom":"FOO"}],"gas":"123456"},"memo":"","msgs":[{"inputs":[{"address":"cosmos1qr90uqx74klw7qx2lcqqz09zqr","coins":[{"amount":"9876543210","denom":"BAR"}]}],"outputs":[{"address":"cosmos1qy352eufqy352euffc88qj","coins":[{"amount":"9876543210","denom":"BAR"}]}]}],"sequence":"42"}'
const unsigned_send_tx_json = {
  "type": "auth/StdTx",
  "value": {
    "msg": [
      {
        "type": "cosmos-sdk/Send",
        "value": {
          "inputs": [
            {
              "address": "cosmos1qr90uqx74klw7qx2lcqqz09zqr",
              "coins": [
                {
                  "denom": "BAR",
                  "amount": "9876543210"
                }
              ]
            }
          ],
          "outputs": [
            {
              "address": "cosmos1qy352eufqy352euffc88qj",
              "coins": [
                {
                  "denom": "BAR",
                  "amount": "9876543210"
                }
              ]
            }
          ]
        }
      }
    ],
    "fee": {
      "amount": [
        {
          "denom": "FOO",
          "amount": "87654321"
        }
      ],
      "gas": "123456"
    },
    "signatures": null,
    "memo": ""
  }
}

const signed_send_tx_binary = 'd101f0625dee0a482a2c87fa0a210a0c00cafe00deadbeef00cafe0012110a03424152120a39383736353433323130121f0a0a0123456789012345678912110a03424152120a3938373635343332313012150a0f0a03464f4f1208383736353433323110c0c4071a6a0a26eb5ae9872103d8290c9a68dc3ed08131148b527b3ac33ff1bf9a03f28232993d101ba827bfbb12404e90693941fff919c07f38528e92b89240cf4a5b79529056b35b87ae3a1340a468c2766ba6b634c6f18b625b904180732dd25433aca95afd8f5c46cbe6bb03b2;'
const signed_send_tx_sign = '{"account_number":"1234","chain_id":"test-chain","fee":{"amount":[{"amount":"87654321","denom":"FOO"}],"gas":"123456"},"memo":"","msgs":[{"inputs":[{"address":"cosmos1qr90uqx74klw7qx2lcqqz09zqr","coins":[{"amount":"9876543210","denom":"BAR"}]}],"outputs":[{"address":"cosmos1qy352eufqy352euffc88qj","coins":[{"amount":"9876543210","denom":"BAR"}]}]}],"sequence":"42"}'
const signed_send_tx_json = {
  "type": "auth/StdTx",
  "value": {
    "msg": [
      {
        "type": "cosmos-sdk/Send",
        "value": {
          "inputs": [
            {
              "address": "cosmos1qr90uqx74klw7qx2lcqqz09zqr",
              "coins": [
                {
                  "denom": "BAR",
                  "amount": "9876543210"
                }
              ]
            }
          ],
          "outputs": [
            {
              "address": "cosmos1qy352eufqy352euffc88qj",
              "coins": [
                {
                  "denom": "BAR",
                  "amount": "9876543210"
                }
              ]
            }
          ]
        }
      }
    ],
    "fee": {
      "amount": [
        {
          "denom": "FOO",
          "amount": "87654321"
        }
      ],
      "gas": "123456"
    },
    "signatures": [
      {
        "pub_key": {
          "type": "tendermint/PubKeySecp256k1",
          "value": "A9gpDJpo3D7QgTEUi1J7OsM/8b+aA/KCMpk9EBuoJ7+7"
        },
        "signature": "TpBpOUH/+RnAfzhSjpK4kkDPSlt5UpBWs1uHrjoTQKRownZrprY0xvGLYluQQYBzLdJUM6ypWv2PXEbL5rsDsg=="
      }
    ],
    "memo": ""
  }
}

