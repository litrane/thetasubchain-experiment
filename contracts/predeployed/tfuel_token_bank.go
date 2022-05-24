package predeployed

import (
	"encoding/hex"
	"math/big"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/crypto"
	"github.com/thetatoken/theta/ledger/types"

	"github.com/thetatoken/thetasubchain/core"
	slst "github.com/thetatoken/thetasubchain/ledger/state"
)

// Bytecode of the smart contracts hardcoded in the genesis block through pre-deployment
const TFuelTokenBankContractBytecode = "608060405234801561001057600080fd5b5060008081905550611ae6806100276000396000f3fe60806040526004361061009c5760003560e01c8063a2cc698111610064578063a2cc6981146101d4578063c27d927a14610211578063d6c7e0d41461023c578063da837d5a14610258578063ebda996214610281578063f6a3d24e146102be5761009c565b80631527b14d146100a1578063261a323e146100df57806327ca4df11461011c578063588b14081461015957806360569b5e14610196575b600080fd5b3480156100ad57600080fd5b506100c860048036038101906100c39190611103565b6102fb565b6040516100d6929190611339565b60405180910390f35b3480156100eb57600080fd5b5061010660048036038101906101019190611103565b610362565b6040516101139190611362565b60405180910390f35b34801561012857600080fd5b50610143600480360381019061013e919061114c565b6103a7565b604051610150919061131e565b60405180910390f35b34801561016557600080fd5b50610180600480360381019061017b919061114c565b6103e6565b60405161018d919061137d565b60405180910390f35b3480156101a257600080fd5b506101bd60048036038101906101b89190611056565b610492565b6040516101cb92919061139f565b60405180910390f35b3480156101e057600080fd5b506101fb60048036038101906101f69190611103565b61054b565b604051610208919061131e565b60405180910390f35b34801561021d57600080fd5b5061022661061b565b604051610233919061146f565b60405180910390f35b61025660048036038101906102519190611083565b610621565b005b34801561026457600080fd5b5061027f600480360381019061027a91906110c3565b6106a5565b005b34801561028d57600080fd5b506102a860048036038101906102a39190611056565b61085c565b6040516102b5919061137d565b60405180910390f35b3480156102ca57600080fd5b506102e560048036038101906102e09190611056565b610989565b6040516102f29190611362565b60405180910390f35b6001818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff16905082565b60008061036e836109e2565b90506001816040516103809190611307565b908152602001604051809103902060000160149054906101000a900460ff16915050919050565b600381815481106103b757600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600481815481106103f657600080fd5b906000526020600020016000915090508054610411906116f8565b80601f016020809104026020016040519081016040528092919081815260200182805461043d906116f8565b801561048a5780601f1061045f5761010080835404028352916020019161048a565b820191906000526020600020905b81548152906001019060200180831161046d57829003601f168201915b505050505081565b60026020528060005260406000206000915090508060000180546104b5906116f8565b80601f01602080910402602001604051908101604052809291908181526020018280546104e1906116f8565b801561052e5780601f106105035761010080835404028352916020019161052e565b820191906000526020600020905b81548152906001019060200180831161051157829003601f168201915b5050505050908060010160009054906101000a900460ff16905082565b600080610557836109e2565b9050600060018260405161056b9190611307565b90815260200160405180910390206040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff161515151581525050905080602001511561060f57806000015192505050610616565b6000925050505b919050565b60005481565b600034905061062f816109f4565b610637610b7b565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f2d8470f382989926e91872b25e27240426cc2efe1bfef95339ae1944a221708f8360005460405161069892919061148a565b60405180910390a3505050565b600080602067ffffffffffffffff8111156106c3576106c2611831565b5b6040519080825280601f01601f1916602001820160405280156106f55781602001600182028036833780820191505090505b50905060008060b573ffffffffffffffffffffffffffffffffffffffff168360405161072191906112f0565b6000604051808303816000865af19150503d806000811461075e576040519150601f19603f3d011682016040523d82523d6000602084013e610763565b606091505b5091509150816107a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079f906113cf565b60405180910390fd5b60006107b382610b96565b9050600181149450846107fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f29061142f565b60405180910390fd5b6108058787610c7f565b8673ffffffffffffffffffffffffffffffffffffffff167fed2a6b76dd30ca61a3c463f15ebbe687c91c02dad5ceea323d771ae5e780e3d18760405161084b919061146f565b60405180910390a250505050505050565b60606000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060400160405290816000820180546108ba906116f8565b80601f01602080910402602001604051908101604052809291908181526020018280546108e6906116f8565b80156109335780601f1061090857610100808354040283529160200191610933565b820191906000526020600020905b81548152906001019060200180831161091657829003601f168201915b505050505081526020016001820160009054906101000a900460ff1615151515815250509050806020015115610970578060000151915050610984565b604051806020016040528060008152509150505b919050565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b60606109ed82610e9e565b9050919050565b6000602067ffffffffffffffff811115610a1157610a10611831565b5b6040519080825280601f01601f191660200182016040528015610a435781602001600182028036833780820191505090505b50905060008260001b905060005b6020811015610ac657818160208110610a6d57610a6c611802565b5b1a60f81b838281518110610a8457610a83611802565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080610abe9061175b565b915050610a51565b50600060b773ffffffffffffffffffffffffffffffffffffffff1683604051610aef91906112f0565b6000604051808303816000865af19150503d8060008114610b2c576040519150601f19603f3d011682016040523d82523d6000602084013e610b31565b606091505b5050905080610b75576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6c9061140f565b60405180910390fd5b50505050565b6001600080828254610b8d9190611546565b92505081905550565b6000806000835190506020811115610be3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bda906113ef565b60405180910390fd5b60005b81811015610c7157600881836020610bfe919061162d565b610c089190611546565b610c1291906115d3565b60ff60f81b868381518110610c2a57610c29611802565b5b602001015160f81c60f81b167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916901c831792508080610c699061175b565b915050610be6565b508160001c92505050919050565b6000603467ffffffffffffffff811115610c9c57610c9b611831565b5b6040519080825280601f01601f191660200182016040528015610cce5781602001600182028036833780820191505090505b50905060008260001b90506000805b6014811015610d61578560601b8160148110610cfc57610cfb611802565b5b1a60f81b848380610d0c9061175b565b945081518110610d1f57610d1e611802565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080610d599061175b565b915050610cdd565b600090505b6020811015610de757828160208110610d8257610d81611802565b5b1a60f81b848380610d929061175b565b945081518110610da557610da4611802565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080610ddf9061175b565b915050610d66565b600060b673ffffffffffffffffffffffffffffffffffffffff1685604051610e0f91906112f0565b6000604051808303816000865af19150503d8060008114610e4c576040519150601f19603f3d011682016040523d82523d6000602084013e610e51565b606091505b5050905080610e95576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e8c9061144f565b60405180910390fd5b50505050505050565b6060600082905060005b8151811015610f2c57610ed7828281518110610ec757610ec6611802565b5b602001015160f81c60f81b610f36565b828281518110610eea57610ee9611802565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080610f249061175b565b915050610ea8565b5080915050919050565b6000604160f81b827effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191610158015610f945750605a60f81b827effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191611155b15610fb35760208260f81c610fa9919061159c565b60f81b9050610fb7565b8190505b919050565b6000610fcf610fca846114d8565b6114b3565b905082815260208101848484011115610feb57610fea611865565b5b610ff68482856116b6565b509392505050565b60008135905061100d81611a82565b92915050565b600082601f83011261102857611027611860565b5b8135611038848260208601610fbc565b91505092915050565b60008135905061105081611a99565b92915050565b60006020828403121561106c5761106b61186f565b5b600061107a84828501610ffe565b91505092915050565b6000806040838503121561109a5761109961186f565b5b60006110a885828601610ffe565b92505060206110b985828601610ffe565b9150509250929050565b600080604083850312156110da576110d961186f565b5b60006110e885828601610ffe565b92505060206110f985828601611041565b9150509250929050565b6000602082840312156111195761111861186f565b5b600082013567ffffffffffffffff8111156111375761113661186a565b5b61114384828501611013565b91505092915050565b6000602082840312156111625761116161186f565b5b600061117084828501611041565b91505092915050565b61118281611661565b82525050565b61119181611673565b82525050565b60006111a282611509565b6111ac818561151f565b93506111bc8185602086016116c5565b80840191505092915050565b60006111d382611514565b6111dd818561152a565b93506111ed8185602086016116c5565b6111f681611874565b840191505092915050565b600061120c82611514565b611216818561153b565b93506112268185602086016116c5565b80840191505092915050565b600061123f604d8361152a565b915061124a82611885565b606082019050919050565b600061126260258361152a565b915061126d826118fa565b604082019050919050565b600061128560488361152a565b915061129082611949565b606082019050919050565b60006112a860388361152a565b91506112b3826119be565b604082019050919050565b60006112cb60478361152a565b91506112d682611a0d565b606082019050919050565b6112ea8161169f565b82525050565b60006112fc8284611197565b915081905092915050565b60006113138284611201565b915081905092915050565b60006020820190506113336000830184611179565b92915050565b600060408201905061134e6000830185611179565b61135b6020830184611188565b9392505050565b60006020820190506113776000830184611188565b92915050565b6000602082019050818103600083015261139781846111c8565b905092915050565b600060408201905081810360008301526113b981856111c8565b90506113c86020830184611188565b9392505050565b600060208201905081810360008301526113e881611232565b9050919050565b6000602082019050818103600083015261140881611255565b9050919050565b6000602082019050818103600083015261142881611278565b9050919050565b600060208201905081810360008301526114488161129b565b9050919050565b60006020820190508181036000830152611468816112be565b9050919050565b600060208201905061148460008301846112e1565b92915050565b600060408201905061149f60008301856112e1565b6114ac60208301846112e1565b9392505050565b60006114bd6114ce565b90506114c9828261172a565b919050565b6000604051905090565b600067ffffffffffffffff8211156114f3576114f2611831565b5b6114fc82611874565b9050602081019050919050565b600081519050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b600081905092915050565b60006115518261169f565b915061155c8361169f565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611591576115906117a4565b5b828201905092915050565b60006115a7826116a9565b91506115b2836116a9565b92508260ff038211156115c8576115c76117a4565b5b828201905092915050565b60006115de8261169f565b91506115e98361169f565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611622576116216117a4565b5b828202905092915050565b60006116388261169f565b91506116438361169f565b925082821015611656576116556117a4565b5b828203905092915050565b600061166c8261167f565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156116e35780820151818401526020810190506116c8565b838111156116f2576000848401525b50505050565b6000600282049050600182168061171057607f821691505b60208210811415611724576117236117d3565b5b50919050565b61173382611874565b810181811067ffffffffffffffff8211171561175257611751611831565b5b80604052505050565b60006117668261169f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611799576117986117a4565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f537562636861696e544675656c546f6b656e42616e6b2e70726976696c65676560008201527f644163636573734f6e6c793a206661696c656420746f20636865636b2074686560208201527f20616363657373206c6576656c00000000000000000000000000000000000000604082015250565b7f627974657320746f2075696e7432353620636f6e76657273696f6e206f76657260008201527f666c6f7773000000000000000000000000000000000000000000000000000000602082015250565b7f537562636861696e544675656c546f6b656e42616e6b2e5f6275726e5446756560008201527f6c566f7563686572733a206661696c656420746f206275726e20544675656c2060208201527f766f756368657273000000000000000000000000000000000000000000000000604082015250565b7f566f75636865724d61702e70726976696c656765644163636573734f6e6c793a60008201527f20696e73756666696369656e742070726976696c656765730000000000000000602082015250565b7f537562636861696e544675656c546f6b656e42616e6b2e6d696e74544675656c60008201527f566f7563686572733a206661696c656420746f206d696e7420544675656c207660208201527f6f75636865727300000000000000000000000000000000000000000000000000604082015250565b611a8b81611661565b8114611a9657600080fd5b50565b611aa28161169f565b8114611aad57600080fd5b5056fea2646970667358221220fb5da469e09786800910b27b1ccad0178ace69ecda6a03cdb81452601bae9b8564736f6c63430008070033"

var mintTFuelVouchersFuncSelector = crypto.Keccak256([]byte("mintVouchers(address,uint256)"))[:4] // In Solidity, uint is an alias of uint256, so we need to use uint256 here to get the correct selector

// TFuelTokenBank implements the TokenBank interface.
type TFuelTokenBank struct {
}

func NewTFuelTokenBank() *TFuelTokenBank {
	return &TFuelTokenBank{}
}

// Mint vouchers for the token transferred cross-chain. If the voucher contract for the token does not yet exist, the
// TokenBank contract deploys the the vouncher contract first and then mints the vouchers in the same call.
// Note: mintVouchers() is only allowed in the privileged execution context. Hence, if a user calls the the TFuelTokenBank.mintVouchers() function (e.g. with the following command),
// the transaction should fail with the "evm revert" error:
//       thetasubcli tx smart_contract --chain="tsub_360777" --from=2E833968E5bB786Ae419c4d13189fB081Cc43bab --to=0xBd770416a3345F91E4B34576cb804a576fa48EB1 --gas_price=4000000000000wei --gas_limit=5000000 --data=da837d5a0000000000000000000000002e833968e5bb786ae419c4d13189fb081cc43bab000000000000000000000000000000000000000000000004c53ecdc18a600000 --password=qwertyuiop --seq=2
func (tb *TFuelTokenBank) GenerateMintVouchersProxySctx(blockProposer common.Address, view *slst.StoreView, ccte *core.CrossChainTFuelTransferEvent) (*types.SmartContractTx, error) {
	voucherReceiver := ccte.Receiver
	amount := ccte.Amount

	calldata := tb.encodeCalldata(voucherReceiver, amount)
	tfuelTokenBankContractAddr := view.GetTFuelTokenBankContractAddress()
	sctx, err := constructProxySmartContractTx(blockProposer, *tfuelTokenBankContractAddr, calldata)
	if err != nil {
		return nil, err
	}

	return sctx, nil
}

// calldata example: da837d5a0000000000000000000000002e833968e5bb786ae419c4d13189fb081cc43bab000000000000000000000000000000000000000000000004c53ecdc18a600000
// Let's break the above calldata into parts, and see what each part represents:
//
// da837d5a // the function selector, i.e mintTFuelVouchersFuncSelector
// 0000000000000000000000002e833968e5bb786ae419c4d13189fb081cc43bab // voucherReceiver, left padded to 32 bytes with zeros
// 000000000000000000000000000000000000000000000004c53ecdc18a600000 // mintAmount, left padded to 32 bytes with zeros, 0x13c9647e25a9940000 = 88000000000000000000L TFuelWei = 99 TFuel
func (tb *TFuelTokenBank) encodeCalldata(voucherReceiver common.Address, amount *big.Int) []byte {
	calldata := append([]byte{}, mintTFuelVouchersFuncSelector...)
	calldata = append(calldata, packAddressParam(voucherReceiver)...)
	calldata = append(calldata, packBigIntParam(amount)...)

	// calldata = append(calldata, common.LeftPadBytes(voucherReceiver.Bytes(), int(wordSizeInBytes))...)
	// calldata = append(calldata, common.LeftPadBytes(packBigIntParam(amount), int(wordSizeInBytes))...)

	logger.Debugf("mint TFuel voucher sctx calldata: %v", hex.EncodeToString(calldata))

	return calldata
}
