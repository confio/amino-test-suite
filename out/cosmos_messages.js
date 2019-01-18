const proposal_binary = '7fb42d614e0a114d792066697273742070726f706f73616c1232546869732069732061206669727374207465737420746f2073656520696620616e796f6e6520766f74657320666f72206d651801220c00cafe00deadbeef00cafe002a0f0a03464f4f120838373635343332312a110a03424152120a39383736353433323130;'
const proposal_json = `{
  "type": "cosmos-sdk/MsgSubmitProposal",
  "value": {
    "title": "My first proposal",
    "description": "This is a first test to see if anyone votes for me",
    "proposal_type": "Text",
    "proposer": "cosmos1qr90uqx74klw7qx2lcqqz09zqr",
    "initial_deposit": [
      {
        "denom": "FOO",
        "amount": "87654321"
      },
      {
        "denom": "BAR",
        "amount": "9876543210"
      }
    ]
  }
}`;

const half_proposal_binary = '19b42d614e1203666f6f1803220c00cafe00deadbeef00cafe00;'
const half_proposal_json = `{
  "type": "cosmos-sdk/MsgSubmitProposal",
  "value": {
    "title": "",
    "description": "foo",
    "proposal_type": "SoftwareUpgrade",
    "proposer": "cosmos1qr90uqx74klw7qx2lcqqz09zqr",
    "initial_deposit": null
  }
}`;

const send_msg_small_binary = '442a2c87fa0a1f0a0c00cafe00deadbeef00cafe00120f0a03464f4f12083837363534333231121d0a0a01234567890123456789120f0a03464f4f12083837363534333231;'
const send_msg_small_json = `{
  "type": "cosmos-sdk/Send",
  "value": {
    "inputs": [
      {
        "address": "cosmos1qr90uqx74klw7qx2lcqqz09zqr",
        "coins": [
          {
            "denom": "FOO",
            "amount": "87654321"
          }
        ]
      }
    ],
    "outputs": [
      {
        "address": "cosmos1qy352eufqy352euffc88qj",
        "coins": [
          {
            "denom": "FOO",
            "amount": "87654321"
          }
        ]
      }
    ]
  }
}`;

const send_msg_swap_binary = '84012a2c87fa0a1d0a0a01234567890123456789120f0a03464f4f120838373635343332310a1f0a0afeedf00d0000be2bee7712110a03424152120a39383736353433323130121f0a0a0123456789012345678912110a03424152120a39383736353433323130121d0a0afeedf00d0000be2bee77120f0a03464f4f12083837363534333231;'
const send_msg_swap_json = `{
  "type": "cosmos-sdk/Send",
  "value": {
    "inputs": [
      {
        "address": "cosmos1qy352eufqy352euffc88qj",
        "coins": [
          {
            "denom": "FOO",
            "amount": "87654321"
          }
        ]
      },
      {
        "address": "cosmos1lmklqrgqqzlzhmnhdtrck5",
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
      },
      {
        "address": "cosmos1lmklqrgqqzlzhmnhdtrck5",
        "coins": [
          {
            "denom": "FOO",
            "amount": "87654321"
          }
        ]
      }
    ]
  }
}`;

const send_msg_big_binary = '8e012a2c87fa0a3d0a0c00cafe00deadbeef00cafe0012110a03424152120a39383736353433323130121a0a03464f4f121331323334353637383930313233343536373839121f0a0a0123456789012345678912110a03424152120a3938373635343332313012280a0afeedf00d0000be2bee77121a0a03464f4f121331323334353637383930313233343536373839;'
const send_msg_big_json = `{
  "type": "cosmos-sdk/Send",
  "value": {
    "inputs": [
      {
        "address": "cosmos1qr90uqx74klw7qx2lcqqz09zqr",
        "coins": [
          {
            "denom": "BAR",
            "amount": "9876543210"
          },
          {
            "denom": "FOO",
            "amount": "1234567890123456789"
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
      },
      {
        "address": "cosmos1lmklqrgqqzlzhmnhdtrck5",
        "coins": [
          {
            "denom": "FOO",
            "amount": "1234567890123456789"
          }
        ]
      }
    ]
  }
}`;

