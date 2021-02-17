package fil

import (
	"encoding/json"
	"testing"
)

func Test_Tipset(t *testing.T) {

	str := `{
   "Cids": [
     {
       "/": "bafy2bzacedrx3uimgtgldpg4sqrml7kaxsnggjusdh6bxt64bpvkvde73mjna"
     },
     {
       "/": "bafy2bzacebs3vzjvfo36x5vvo63pbglyajzsob2qfd5chdgi3foge2mgkemzw"
     }
   ],
   "Blocks": [
     {
       "Miner": "t01002",
       "Ticket": {
         "VRFProof": "g+bM5TkO3yQZro+jeEYWnaafOTJW1aMmvu0eKsSB+zC9mUYOO6SL7W6JiKCsi3zpE5uH/GUJ3y693yfzWYpcaNI8uRC2xpz2emjfN3Lr2AfAXaf5ov42CGCAoqYvudTJ"
       },
       "ElectionProof": {
         "VRFProof": "tBIMq/VIOy+YeVMg9t8rHplu4GkEQkm0OilLQ9B0U8Vaw5f1Ei+ixgbHBpKy+j6yBL5P5vvAGPyhWfyFLhKW7a7yb8tQWLkl8zDaemvk6sfoMxwus95cIi7GhhUItW4s"
       },
       "BeaconEntries": null,
       "WinPoStProof": [
         {
           "RegisteredProof": 15,
           "ProofBytes": "laXXoDzQE/cb0UrjAVWYdHtGdynyJAG9dZmmsIvGjpjpF6iUBmu3OSDu10s+5jbrkkSS31PDurEARweHMDEHeu/lEfyx+QVR/arbwD+/XSSWolrvSpG+wxt/eIKdNTnzCsLZia5fhwkibaVi53KOjafsplgnCZy6pH3gDIW94fljwddo3j8UKYb8XVEroHtBoVWSLK8L6vGfqbcYv/mV1jitfRkJBFy4afNx/rcobG3SJFHSMq9gtE8Nv5JoPbm7"
         }
       ],
       "Parents": [
         {
           "/": "bafy2bzaceci6m7zlvrip3ufhgotnc6t3tsc3qj6avz5mvukrvs5a6v4qycevk"
         },
         {
           "/": "bafy2bzacece62mvuc4v7qwyhicb3cqzoucbi5wxdx2vqc6nutylbou5xlj6p4"
         }
       ],
       "ParentWeight": "25964240",
       "Height": 2000,
       "ParentStateRoot": {
         "/": "bafy2bzaceamrgqt5cqz2up2eb2ttdtjd5fzqsdyj32utool45egcrpfmicfna"
       },
       "ParentMessageReceipts": {
         "/": "bafy2bzacecbm7dxxaz36pcfbam4p4bzffksejj72zeyrltu4zt46ztcf52qy6"
       },
       "Messages": {
         "/": "bafy2bzacebjqxrrfmqzwboiilvot55gzyyfzujbuulks7s2dtolbvazuwgso4"
       },
       "BLSAggregate": {
         "Type": 2,
         "Data": "owcZhdrk+H95CQMdeUGJSPo+7IeaPm8BmSJ+X7fs7VzJRxgoupW4pfyCUYlIUDXrBYb0VSLPPcJ8aihQ63hCwSxM7eYag2m0vPLHoZUBTYJUZP1lx2nafulvsvmBgjPL"
       },
       "Timestamp": 1589544200,
       "BlockSig": {
         "Type": 2,
         "Data": "hi6fpNcX9LupN4HUG+FojvVT0KseEMLHx+udT0shrAWS3DhpeSdh2yypowB/V1eHE2ti4/Jwquqjw+5fS7mB7B1HTlqsiR3WVs43nh+8emZm2mXjlidsrOklZqJZiFxw"
       },
       "ForkSignaling": 0
     },
     {
       "Miner": "t01001",
       "Ticket": {
         "VRFProof": "oMUhqYwENAnudzAIH9ZYyOB+6tPjDoHV7A+q4BlH+6ei+4JDniky7cMZGFqBmaJfFDEA1+olOIHhLcoEyWoe+GxlQHoiTY6S+kkUiR8xcuusgMOBMgSvNhCbqO+igbZd"
       },
       "ElectionProof": {
         "VRFProof": "uHCpet8zbhFTak7Iw4Z4eZANufO4o/8pAYLkcYa+0GtmtgapZWB6R6+uUJSP/1nDEvd2E0SmwSdsPGWwUDicdB08mgHLHcnJVvmQY0OSogWECOTyqUWrh75ovZeVAfFQ"
       },
       "BeaconEntries": null,
       "WinPoStProof": [
         {
           "RegisteredProof": 15,
           "ProofBytes": "qx7FsgQ7x+B1DmwcWrB8VD0z4fAJGlS92bhLO+/PdgdZdikDgLbWM2ehd17dJZVGhwLEhcAiDvxDOvLMHPvAl1dmSFYxVho5EYqjv7hY4ns/uWdiS7cS0eBO8EaVZ0f4AiCQmRmjAeEyyuwhtO+1qfYq2ZHV4gyFFm/sEMOpvv/P53SqWi/cxkirkxVmI3qojaZrvL7rwyEhalrbbN/WV3s+gx4lkh812oxgRggMyxXbnldcLJx/77mOyAKVCuiR"
         }
       ],
       "Parents": [
         {
           "/": "bafy2bzaceci6m7zlvrip3ufhgotnc6t3tsc3qj6avz5mvukrvs5a6v4qycevk"
         },
         {
           "/": "bafy2bzacece62mvuc4v7qwyhicb3cqzoucbi5wxdx2vqc6nutylbou5xlj6p4"
         }
       ],
       "ParentWeight": "25964240",
       "Height": 2000,
       "ParentStateRoot": {
         "/": "bafy2bzaceamrgqt5cqz2up2eb2ttdtjd5fzqsdyj32utool45egcrpfmicfna"
       },
       "ParentMessageReceipts": {
         "/": "bafy2bzacecbm7dxxaz36pcfbam4p4bzffksejj72zeyrltu4zt46ztcf52qy6"
       },
       "Messages": {
         "/": "bafy2bzacecbgklp6rhvecjtfnwg2hcu3skuegcfi3zjur6c7cwep2y2oaqim4"
       },
       "BLSAggregate": {
         "Type": 2,
         "Data": "owcZhdrk+H95CQMdeUGJSPo+7IeaPm8BmSJ+X7fs7VzJRxgoupW4pfyCUYlIUDXrBYb0VSLPPcJ8aihQ63hCwSxM7eYag2m0vPLHoZUBTYJUZP1lx2nafulvsvmBgjPL"
       },
       "Timestamp": 1589544200,
       "BlockSig": {
         "Type": 2,
         "Data": "h34neNK9asLMmG7VPmJMN67Dr/IvHKV78YTuf4tfKDtLA9a43wMSp4oZhpYP1BD7CCv/Pkl8rK/bWhksDgPoHY7cMcFK0xzz5aBTTYxcODcOII+ruqedpSamjniDemUS"
       },
       "ForkSignaling": 0
     }
   ],
   "Height": 2000
  }`

	var tet = new(Tipset)
	json.Unmarshal([]byte(str), &tet)

	if tet.Cids[0].Blockcid != "bafy2bzacedrx3uimgtgldpg4sqrml7kaxsnggjusdh6bxt64bpvkvde73mjna" {
		t.Fatal("Test_ReportDepositForm fail ")
	}

}
