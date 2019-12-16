package owaddress

import (
	"github.com/Assetsadapter/go-owaddress/coins/ae"
	"github.com/Assetsadapter/go-owaddress/coins/alc"
	"github.com/Assetsadapter/go-owaddress/coins/amtc"
	"github.com/Assetsadapter/go-owaddress/coins/atom"
	"github.com/Assetsadapter/go-owaddress/coins/bnb"
	"github.com/Assetsadapter/go-owaddress/coins/btc"
	"github.com/Assetsadapter/go-owaddress/coins/btx"
	"github.com/Assetsadapter/go-owaddress/coins/cxc"
	"github.com/Assetsadapter/go-owaddress/coins/dgb"
	"github.com/Assetsadapter/go-owaddress/coins/dsc"
	"github.com/Assetsadapter/go-owaddress/coins/ela"
	"github.com/Assetsadapter/go-owaddress/coins/eth"
	"github.com/Assetsadapter/go-owaddress/coins/etz"
	"github.com/Assetsadapter/go-owaddress/coins/fiii"
	"github.com/Assetsadapter/go-owaddress/coins/g50"
	"github.com/Assetsadapter/go-owaddress/coins/hc"
	"github.com/Assetsadapter/go-owaddress/coins/hss"
	"github.com/Assetsadapter/go-owaddress/coins/ltc"
	"github.com/Assetsadapter/go-owaddress/coins/macc"
	"github.com/Assetsadapter/go-owaddress/coins/moac"
	"github.com/Assetsadapter/go-owaddress/coins/nas"
	"github.com/Assetsadapter/go-owaddress/coins/neo"
	"github.com/Assetsadapter/go-owaddress/coins/ont"
	"github.com/Assetsadapter/go-owaddress/coins/pess"
	"github.com/Assetsadapter/go-owaddress/coins/qtum"
	"github.com/Assetsadapter/go-owaddress/coins/ria"
	"github.com/Assetsadapter/go-owaddress/coins/sinoc"
	"github.com/Assetsadapter/go-owaddress/coins/spg"
	truechain "github.com/Assetsadapter/go-owaddress/coins/true"
	"github.com/Assetsadapter/go-owaddress/coins/trx"
	"github.com/Assetsadapter/go-owaddress/coins/tv"
	"github.com/Assetsadapter/go-owaddress/coins/vcc"
	"github.com/Assetsadapter/go-owaddress/coins/vds"
	"github.com/Assetsadapter/go-owaddress/coins/vlx"
	"github.com/Assetsadapter/go-owaddress/coins/vns"
	"github.com/Assetsadapter/go-owaddress/coins/vsys"
	"github.com/Assetsadapter/go-owaddress/coins/wicc"
	"github.com/Assetsadapter/go-owaddress/coins/xrp"
	"github.com/Assetsadapter/go-owaddress/coins/xvg"
	"github.com/Assetsadapter/go-owaddress/coins/zen"
	"reflect"
)

var AddressVerifyRegistry = make(map[string]reflect.Type)

func RegisterAddressVerify(elem interface{}, coin string) {
	t := reflect.TypeOf(elem).Elem()
	AddressVerifyRegistry[coin] = t
}

func init() {
	RegisterAddressVerify(btc.DefaultStruct, btc.CoinName)
	RegisterAddressVerify(ltc.DefaultStruct, ltc.CoinName)
	RegisterAddressVerify(qtum.DefaultStruct, qtum.CoinName)
	RegisterAddressVerify(ont.DefaultStruct, ont.CoinName)
	RegisterAddressVerify(atom.DefaultStruct, atom.CoinName)
	RegisterAddressVerify(xrp.DefaultStruct, xrp.CoinName)
	RegisterAddressVerify(trx.DefaultStruct, trx.CoinName)
	RegisterAddressVerify(wicc.DefaultStruct, wicc.CoinName)
	RegisterAddressVerify(hc.DefaultStruct, hc.CoinName)
	RegisterAddressVerify(bnb.DefaultStruct, bnb.CoinName)
	RegisterAddressVerify(vsys.DefaultStruct, vsys.CoinName)
	RegisterAddressVerify(nas.DefaultStruct, nas.CoinName)
	RegisterAddressVerify(ela.DefaultStruct, ela.CoinName)
	RegisterAddressVerify(tv.DefaultStruct, tv.CoinName)
	RegisterAddressVerify(moac.DefaultStruct, moac.CoinName)
	RegisterAddressVerify(dsc.DefaultStruct, dsc.CoinName)
	RegisterAddressVerify(fiii.DefaultStruct, fiii.CoinName)
	RegisterAddressVerify(vds.DefaultStruct, vds.CoinName)
	RegisterAddressVerify(hss.DefaultStruct, hss.CoinName)
	RegisterAddressVerify(vlx.DefaultStruct, vlx.CoinName)
	RegisterAddressVerify(btx.DefaultStruct, btx.CoinName)
	RegisterAddressVerify(cxc.DefaultStruct, cxc.CoinName)
	RegisterAddressVerify(xvg.DefaultStruct, xvg.CoinName)
	RegisterAddressVerify(zen.DefaultStruct, zen.CoinName)
	RegisterAddressVerify(dgb.DefaultStruct, dgb.CoinName)
	RegisterAddressVerify(alc.DefaultStruct, alc.CoinName)
	RegisterAddressVerify(eth.DefaultStruct, eth.CoinName)
	RegisterAddressVerify(etz.DefaultStruct, etz.CoinName)
	RegisterAddressVerify(pess.DefaultStruct, pess.CoinName)
	RegisterAddressVerify(vcc.DefaultStruct, vcc.CoinName)
	RegisterAddressVerify(truechain.DefaultStruct, truechain.CoinName)
	RegisterAddressVerify(g50.DefaultStruct, g50.CoinName)
	RegisterAddressVerify(sinoc.DefaultStruct, sinoc.CoinName)
	RegisterAddressVerify(ae.DefaultStruct, ae.CoinName)
	RegisterAddressVerify(macc.DefaultStruct, macc.CoinName)
	RegisterAddressVerify(amtc.DefaultStruct, amtc.CoinName)
	RegisterAddressVerify(vns.DefaultStruct, vns.CoinName)
	RegisterAddressVerify(spg.DefaultStruct, spg.CoinName)
	RegisterAddressVerify(neo.DefaultStruct, neo.CoinName)
	RegisterAddressVerify(ria.DefaultStruct, ria.CoinName)
}
