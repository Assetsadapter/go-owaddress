package owaddress

import (
	"reflect"

	"github.com/Assetsadapter/go-owaddress/coins/ae"
	"github.com/Assetsadapter/go-owaddress/coins/alc"
	"github.com/Assetsadapter/go-owaddress/coins/ark"
	"github.com/Assetsadapter/go-owaddress/coins/atom"
	"github.com/Assetsadapter/go-owaddress/coins/avax"
	"github.com/Assetsadapter/go-owaddress/coins/bbc"
	"github.com/Assetsadapter/go-owaddress/coins/bch"
	"github.com/Assetsadapter/go-owaddress/coins/beth"
	"github.com/Assetsadapter/go-owaddress/coins/bnb"
	"github.com/Assetsadapter/go-owaddress/coins/bsv"
	"github.com/Assetsadapter/go-owaddress/coins/btc"
	"github.com/Assetsadapter/go-owaddress/coins/btx"
	"github.com/Assetsadapter/go-owaddress/coins/cxc"
	"github.com/Assetsadapter/go-owaddress/coins/dgb"
	"github.com/Assetsadapter/go-owaddress/coins/dot"
	"github.com/Assetsadapter/go-owaddress/coins/dsc"
	"github.com/Assetsadapter/go-owaddress/coins/ela"
	"github.com/Assetsadapter/go-owaddress/coins/eth"
	"github.com/Assetsadapter/go-owaddress/coins/etp"
	"github.com/Assetsadapter/go-owaddress/coins/eva"
	"github.com/Assetsadapter/go-owaddress/coins/fac"
	"github.com/Assetsadapter/go-owaddress/coins/fiii"
	"github.com/Assetsadapter/go-owaddress/coins/fil"
	"github.com/Assetsadapter/go-owaddress/coins/g50"
	"github.com/Assetsadapter/go-owaddress/coins/hc"
	"github.com/Assetsadapter/go-owaddress/coins/hns"
	"github.com/Assetsadapter/go-owaddress/coins/hss"
	"github.com/Assetsadapter/go-owaddress/coins/ilc"
	"github.com/Assetsadapter/go-owaddress/coins/kpg"
	"github.com/Assetsadapter/go-owaddress/coins/ksm"
	"github.com/Assetsadapter/go-owaddress/coins/ltc"
	"github.com/Assetsadapter/go-owaddress/coins/macc"
	"github.com/Assetsadapter/go-owaddress/coins/moac"
	"github.com/Assetsadapter/go-owaddress/coins/nas"
	"github.com/Assetsadapter/go-owaddress/coins/nhss"
	"github.com/Assetsadapter/go-owaddress/coins/ntn"
	"github.com/Assetsadapter/go-owaddress/coins/nuls2"
	"github.com/Assetsadapter/go-owaddress/coins/ont"
	"github.com/Assetsadapter/go-owaddress/coins/pb"
	"github.com/Assetsadapter/go-owaddress/coins/pess"
	"github.com/Assetsadapter/go-owaddress/coins/qtum"
	"github.com/Assetsadapter/go-owaddress/coins/rcp"
	"github.com/Assetsadapter/go-owaddress/coins/sgu"
	"github.com/Assetsadapter/go-owaddress/coins/sinoc"
	truechain "github.com/Assetsadapter/go-owaddress/coins/true"
	"github.com/Assetsadapter/go-owaddress/coins/trx"
	"github.com/Assetsadapter/go-owaddress/coins/tv"
	"github.com/Assetsadapter/go-owaddress/coins/ufc"
	"github.com/Assetsadapter/go-owaddress/coins/vas"
	"github.com/Assetsadapter/go-owaddress/coins/vcc"
	"github.com/Assetsadapter/go-owaddress/coins/vds"
	"github.com/Assetsadapter/go-owaddress/coins/vlx"
	"github.com/Assetsadapter/go-owaddress/coins/vsys"
	"github.com/Assetsadapter/go-owaddress/coins/wicc"
	"github.com/Assetsadapter/go-owaddress/coins/xif"
	"github.com/Assetsadapter/go-owaddress/coins/xrp"
	"github.com/Assetsadapter/go-owaddress/coins/xvg"
	"github.com/Assetsadapter/go-owaddress/coins/xwc"
	"github.com/Assetsadapter/go-owaddress/coins/zen"
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
	RegisterAddressVerify(dot.DefaultStruct, dot.CoinName)
	RegisterAddressVerify(fil.DefaultStruct, fil.CoinName)
	RegisterAddressVerify(avax.DefaultStruct, avax.CoinName)
	RegisterAddressVerify(ksm.DefaultStruct, ksm.CoinName)
	RegisterAddressVerify(ufc.DefaultStruct, ufc.CoinName)
}
