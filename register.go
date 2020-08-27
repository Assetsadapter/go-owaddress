package owaddress

import (
	"github.com/Assetsadapter/go-owaddress/v2/coins/ae"
	"github.com/Assetsadapter/go-owaddress/v2/coins/alc"
	"github.com/Assetsadapter/go-owaddress/v2/coins/ark"
	"github.com/Assetsadapter/go-owaddress/v2/coins/atom"
	"github.com/Assetsadapter/go-owaddress/v2/coins/bbc"
	"github.com/Assetsadapter/go-owaddress/v2/coins/bch"
	"github.com/Assetsadapter/go-owaddress/v2/coins/beth"
	"github.com/Assetsadapter/go-owaddress/v2/coins/bnb"
	"github.com/Assetsadapter/go-owaddress/v2/coins/bsv"
	"github.com/Assetsadapter/go-owaddress/v2/coins/btc"
	"github.com/Assetsadapter/go-owaddress/v2/coins/btx"
	"github.com/Assetsadapter/go-owaddress/v2/coins/cxc"
	"github.com/Assetsadapter/go-owaddress/v2/coins/dgb"
	"github.com/Assetsadapter/go-owaddress/v2/coins/dot"
	"github.com/Assetsadapter/go-owaddress/v2/coins/dsc"
	"github.com/Assetsadapter/go-owaddress/v2/coins/ela"
	"github.com/Assetsadapter/go-owaddress/v2/coins/eth"
	"github.com/Assetsadapter/go-owaddress/v2/coins/etp"
	"github.com/Assetsadapter/go-owaddress/v2/coins/eva"
	"github.com/Assetsadapter/go-owaddress/v2/coins/fiii"
	"github.com/Assetsadapter/go-owaddress/v2/coins/g50"
	"github.com/Assetsadapter/go-owaddress/v2/coins/hc"
	"github.com/Assetsadapter/go-owaddress/v2/coins/hns"
	"github.com/Assetsadapter/go-owaddress/v2/coins/hss"
	"github.com/Assetsadapter/go-owaddress/v2/coins/ilc"
	"github.com/Assetsadapter/go-owaddress/v2/coins/kpg"
	"github.com/Assetsadapter/go-owaddress/v2/coins/ltc"
	"github.com/Assetsadapter/go-owaddress/v2/coins/macc"
	"github.com/Assetsadapter/go-owaddress/v2/coins/moac"
	"github.com/Assetsadapter/go-owaddress/v2/coins/nas"
	"github.com/Assetsadapter/go-owaddress/v2/coins/ntn"
	"github.com/Assetsadapter/go-owaddress/v2/coins/nuls2"
	"github.com/Assetsadapter/go-owaddress/v2/coins/ont"
	"github.com/Assetsadapter/go-owaddress/v2/coins/pb"
	"github.com/Assetsadapter/go-owaddress/v2/coins/pess"
	"github.com/Assetsadapter/go-owaddress/v2/coins/qtum"
	"github.com/Assetsadapter/go-owaddress/v2/coins/rcp"
	"github.com/Assetsadapter/go-owaddress/v2/coins/sgu"
	"github.com/Assetsadapter/go-owaddress/v2/coins/sinoc"
	truechain "github.com/Assetsadapter/go-owaddress/v2/coins/true"
	"github.com/Assetsadapter/go-owaddress/v2/coins/trx"
	"github.com/Assetsadapter/go-owaddress/v2/coins/tv"
	"github.com/Assetsadapter/go-owaddress/v2/coins/vas"
	"github.com/Assetsadapter/go-owaddress/v2/coins/vcc"
	"github.com/Assetsadapter/go-owaddress/v2/coins/vds"
	"github.com/Assetsadapter/go-owaddress/v2/coins/vlx"
	"github.com/Assetsadapter/go-owaddress/v2/coins/vsys"
	"github.com/Assetsadapter/go-owaddress/v2/coins/wicc"
	"github.com/Assetsadapter/go-owaddress/v2/coins/xif"
	"github.com/Assetsadapter/go-owaddress/v2/coins/xrp"
	"github.com/Assetsadapter/go-owaddress/v2/coins/fac"
	"github.com/Assetsadapter/go-owaddress/v2/coins/xvg"
	"github.com/Assetsadapter/go-owaddress/v2/coins/xwc"
	"github.com/Assetsadapter/go-owaddress/v2/coins/zen"
	"github.com/Assetsadapter/go-owaddress/v2/coins/nhss"
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
	RegisterAddressVerify(pess.DefaultStruct, pess.CoinName)
	RegisterAddressVerify(vcc.DefaultStruct, vcc.CoinName)
	RegisterAddressVerify(truechain.DefaultStruct, truechain.CoinName)
	RegisterAddressVerify(g50.DefaultStruct, g50.CoinName)
	RegisterAddressVerify(sinoc.DefaultStruct, sinoc.CoinName)
	RegisterAddressVerify(ae.DefaultStruct, ae.CoinName)
	RegisterAddressVerify(macc.DefaultStruct, macc.CoinName)
	RegisterAddressVerify(ntn.DefaultStruct, ntn.CoinName)
	RegisterAddressVerify(beth.DefaultStruct, beth.CoinName)
	RegisterAddressVerify(xwc.DefaultStruct, xwc.CoinName)
	RegisterAddressVerify(vas.DefaultStruct, vas.CoinName)
	RegisterAddressVerify(etp.DefaultStruct, etp.CoinName)
	RegisterAddressVerify(nuls2.DefaultStruct, nuls2.CoinName)
	RegisterAddressVerify(bch.DefaultStruct, bch.CoinName)
	RegisterAddressVerify(bsv.DefaultStruct, bsv.CoinName)
	RegisterAddressVerify(bbc.DefaultStruct, bbc.CoinName)
	RegisterAddressVerify(sgu.DefaultStruct, sgu.CoinName)
	RegisterAddressVerify(eva.DefaultStruct, eva.CoinName)
	RegisterAddressVerify(rcp.DefaultStruct, rcp.CoinName)
	RegisterAddressVerify(pb.DefaultStruct, pb.CoinName)
	RegisterAddressVerify(ark.DefaultStruct, ark.CoinName)
	RegisterAddressVerify(fac.DefaultStruct, fac.CoinName)
	RegisterAddressVerify(nhss.DefaultStruct, nhss.CoinName)
	RegisterAddressVerify(ilc.DefaultStruct, ilc.CoinName)
	RegisterAddressVerify(xif.DefaultStruct, xif.CoinNameXIF)
	RegisterAddressVerify(xif.DefaultStruct, xif.CoinNameAUSD)
	RegisterAddressVerify(hns.DefaultStruct, hns.CoinName)
	RegisterAddressVerify(kpg.DefaultStruct, kpg.CoinName)
	RegisterAddressVerify(dot.DefaultStruct, dot.CoinNameDOT)
	RegisterAddressVerify(dot.DefaultStruct, dot.CoinNameKSM)
}
