package alipay

import (
	"github.com/yaooort/gopay"
	"github.com/yaooort/gopay/alipay"
	"github.com/go-pay/xlog"
)

func TradeCreate() {
	//aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"
	//privateKey := "MIIEpgIBAAKCAQEA5JwIRwps/xrBFLCRHL727tofHEDE74I+NMkUkHCEd0qI7lOYrz72761hjw0/Np+io8DQDNf3AV0+HJGQAuMadhgmcLUnfB7BJwEL1wgc2Uk8oj8XBIJUKmc+APtJ3NhUgd7s4p63IVKjQfp8oMbGZS3/dJdamBEqGJn3bep7WlKyJy5bL3KY+fX5F1DNif8xfHBSygkniKhTt/pkKaEP6uiEvhZuorPKGTowU99cqJZAqNVGcZV/30UEerR9BHV5g3iZHRxebGNbC9WSBqVWoGhdo23Xl0QG6KmuZ1akF7mdq1e7Qlpm0SkYimqipRhSflu2N/mIW2pesypX2DRE8wIDAQABAoIBAQDDvkJsKaYoTI00nPaziWPAIm3u8BpwBbIj2mTMRDt4NQ7IGjYrH5uqn2dgaHRO3iMRWYAK70RlG6SaK50Gi9RVd9o6OTKk2WSDdCbiKOUiu826EpEv4DQW1q3Fg2S8P1MknG9yn3mog6ycStE5jNnFOrx35TzB5jiINhUhPVv/eh6iDkHCxfOh2R4+vPYjNXlzhgcR3qh19Rnd2uuzscbRtljoZpA3qE0tIawW7HmesKqjfHL3O/kHJGmZjK/uhNvZmPMzilmFb+q3R3BXbAPNc/efkv0GF1fS0Xy2h0YYfkNTDzewsU7Qwkw/iT29IUUn9uxAux6SYyubCdLRIMtpAoGBAP3uxgt0kX4AN0CUibkyvDlaw+amkjTgk9voH78CPUej58sIkWcaPJJwAFOhzNbsTlNp3mZ2xx58gqJU2Rm1/6bsQbYRjN9qVY8Fco7ClysOaf9bwDBqnIjv9eHZdw8EmR4qnPCH9yjBWDMODKxIhAhsuwm4D7GmsHrFAp5w7PlnAoGBAOZ4e2sOfVAiM9ZoHIlaNLW2Z1PKkc/1k+Hxoxl5dq9nHB9RAGqpO4vjbMc6d1wVRr02GZE8RaLSzS7WrAgSzsWgtql4P8TvrjPvIPEyN1B7/eqR6yNDGRHLkYbDg54SeRI/yUg7EwMARJqa24zP503IA40JZgaVHLti5XYCIoSVAoGBAIx/oXpTG/EQY4HK7czXgodlbgDfZwP0wjqpW29O6OMLkEOpUPIv9RW8/KFJY9IC22+RoykkTRXUJF92/MwaBAKys8TuPviamA5TIEhW6Fc9WW6dsF/ZjRTDWFOHBDX+AE5Nm5oGUL1vBMLy4hYs5UjYCEDfY2eS6BB1kvZhWuy/AoGBAKJHf4+YOkBsdNepoz5LmAXDE+p6HkWyA5j6jf9n+Vv0XGbooER6OQV886Es19Ks4IxmRYZwTBAkInmyipt8sr+RNE0L8Mr4gU7sN+Pdmfk/9UBv0oOXwVU4Y5XioRByrXFeHJqRjgd07tl15NW3poSsK2PplD9aS0rejfeT9T4pAoGBAKELHKfRAWILgkI/L5RrvaoUEkdy1hw9a+b+V1TKB+wDbrsomzL2BPnBYjqtw2ZhTffpxzS3ZpSzlMtgmcYhpyAvq5GlPw+wbsNXr4tCIu8eYLc6QoypmO8kE2FornyZfol5ygEu6HSUSwEHHjIQBDxtAlYgBPZOetDVdDVJD1nR"
	privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境
	client, err := alipay.NewClient("2016091200494382", privateKey, false)
	if err != nil {
		xlog.Error(err)
		return
	}
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		SetNotifyUrl("https://www.fmm.ink")

	//请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "创建订单").
		Set("buyer_id", "2088802095984694").
		Set("out_trade_no", "GZ201901301040355709").
		Set("total_amount", "0.01")
	//创建订单
	aliRsp, err := client.TradeCreate(ctx, bm)
	if err != nil {
		xlog.Error("err:", err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
	xlog.Debug("aliRsp.TradeNo:", aliRsp.Response.TradeNo)
}
