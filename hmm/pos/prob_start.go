package pos

var probStart = map[uint16]float64{
	100: -4.762305214596967,
	101: -6.680066036784177,
	102: -3.14e+100,
	103: -8.697083223018778,
	104: -5.018374362109218,
	105: -3.14e+100,
	106: -3.423880184954888,
	107: -3.9750475297585357,
	108: -8.888974230828882,
	109: -3.14e+100,
	110: -8.563551830394255,
	111: -3.14e+100,
	112: -5.491630418482717,
	113: -3.14e+100,
	114: -13.533365129970255,
	115: -6.1157847275557105,
	116: -3.14e+100,
	117: -5.0576191284681915,
	118: -3.14e+100,
	119: -3.14e+100,
	120: -4.905883584659895,
	121: -3.14e+100,
	122: -3.6524299819046386,
	123: -3.14e+100,
	124: -6.78695300139688,
	125: -1.6966257797548328,
	126: -3.14e+100,
	127: -2.2310495913769506,
	128: -5.873722175405573,
	129: -4.985642733519195,
	130: -2.8228438314969213,
	131: -4.846091668182416,
	132: -3.94698846057672,
	133: -8.433498702146057,
	134: -4.200984132085048,
	135: -6.998123858956596,
	136: -3.14e+100,
	137: -3.14e+100,
	138: -3.4098187790818413,
	139: -3.14e+100,
	140: -12.434752841302146,
	141: -7.946116471570005,
	142: -5.522673590839954,
	143: -3.3647479094528574,
	144: -3.14e+100,
	145: -9.163917277503234,
	146: -3.14e+100,
	147: -3.14e+100,
	148: -3.14e+100,
	149: -3.14e+100,
	150: -3.14e+100,
	151: -3.14e+100,
	152: -2.6740584874265685,
	153: -9.044728760238115,
	154: -3.14e+100,
	155: -12.434752841302146,
	156: -4.3315610890163585,
	157: -12.147070768850364,
	158: -3.14e+100,
	159: -3.14e+100,
	160: -9.844485675856319,
	161: -3.14e+100,
	162: -7.045681111485645,
	163: -3.14e+100,
	200: -3.14e+100,
	201: -3.14e+100,
	202: -3.14e+100,
	203: -3.14e+100,
	204: -3.14e+100,
	205: -3.14e+100,
	206: -3.14e+100,
	207: -3.14e+100,
	208: -3.14e+100,
	209: -3.14e+100,
	210: -3.14e+100,
	211: -3.14e+100,
	212: -3.14e+100,
	213: -3.14e+100,
	214: -3.14e+100,
	215: -3.14e+100,
	216: -3.14e+100,
	217: -3.14e+100,
	218: -3.14e+100,
	219: -3.14e+100,
	220: -3.14e+100,
	221: -3.14e+100,
	222: -3.14e+100,
	223: -3.14e+100,
	224: -3.14e+100,
	225: -3.14e+100,
	226: -3.14e+100,
	227: -3.14e+100,
	228: -3.14e+100,
	229: -3.14e+100,
	230: -3.14e+100,
	231: -3.14e+100,
	232: -3.14e+100,
	233: -3.14e+100,
	234: -3.14e+100,
	235: -3.14e+100,
	236: -3.14e+100,
	237: -3.14e+100,
	238: -3.14e+100,
	239: -3.14e+100,
	240: -3.14e+100,
	241: -3.14e+100,
	242: -3.14e+100,
	243: -3.14e+100,
	244: -3.14e+100,
	245: -3.14e+100,
	246: -3.14e+100,
	247: -3.14e+100,
	248: -3.14e+100,
	249: -3.14e+100,
	250: -3.14e+100,
	251: -3.14e+100,
	252: -3.14e+100,
	253: -3.14e+100,
	254: -3.14e+100,
	255: -3.14e+100,
	256: -3.14e+100,
	257: -3.14e+100,
	258: -3.14e+100,
	259: -3.14e+100,
	260: -3.14e+100,
	261: -3.14e+100,
	262: -3.14e+100,
	263: -3.14e+100,
	300: -3.14e+100,
	301: -3.14e+100,
	302: -3.14e+100,
	303: -3.14e+100,
	304: -3.14e+100,
	305: -3.14e+100,
	306: -3.14e+100,
	307: -3.14e+100,
	308: -3.14e+100,
	309: -3.14e+100,
	310: -3.14e+100,
	311: -3.14e+100,
	312: -3.14e+100,
	313: -3.14e+100,
	314: -3.14e+100,
	315: -3.14e+100,
	316: -3.14e+100,
	317: -3.14e+100,
	318: -3.14e+100,
	319: -3.14e+100,
	320: -3.14e+100,
	321: -3.14e+100,
	322: -3.14e+100,
	323: -3.14e+100,
	324: -3.14e+100,
	325: -3.14e+100,
	326: -3.14e+100,
	327: -3.14e+100,
	328: -3.14e+100,
	329: -3.14e+100,
	330: -3.14e+100,
	331: -3.14e+100,
	332: -3.14e+100,
	333: -3.14e+100,
	334: -3.14e+100,
	335: -3.14e+100,
	336: -3.14e+100,
	337: -3.14e+100,
	338: -3.14e+100,
	339: -3.14e+100,
	340: -3.14e+100,
	341: -3.14e+100,
	342: -3.14e+100,
	343: -3.14e+100,
	344: -3.14e+100,
	345: -3.14e+100,
	346: -3.14e+100,
	347: -3.14e+100,
	348: -3.14e+100,
	349: -3.14e+100,
	350: -3.14e+100,
	351: -3.14e+100,
	352: -3.14e+100,
	353: -3.14e+100,
	354: -3.14e+100,
	355: -3.14e+100,
	356: -3.14e+100,
	357: -3.14e+100,
	358: -3.14e+100,
	359: -3.14e+100,
	360: -3.14e+100,
	361: -3.14e+100,
	362: -3.14e+100,
	363: -3.14e+100,
	400: -3.9025396831295227,
	401: -11.048458480182255,
	402: -6.954113917960154,
	403: -12.84021794941031,
	404: -6.472888763970454,
	405: -3.14e+100,
	406: -4.786966795861212,
	407: -3.903919764181873,
	408: -3.14e+100,
	409: -8.948397651299683,
	410: -5.942513006281674,
	411: -3.14e+100,
	412: -5.194820249981676,
	413: -6.507826815331734,
	414: -8.650563207383884,
	415: -3.14e+100,
	416: -3.14e+100,
	417: -4.911992119644354,
	418: -3.14e+100,
	419: -6.940320595827818,
	420: -3.14e+100,
	421: -3.14e+100,
	422: -3.269200652116097,
	423: -10.825314928868044,
	424: -3.14e+100,
	425: -3.8551483897645107,
	426: -4.913434861102905,
	427: -4.483663103956885,
	428: -3.14e+100,
	429: -3.14e+100,
	430: -3.14e+100,
	431: -12.147070768850364,
	432: -3.14e+100,
	433: -8.464460927750023,
	434: -2.9868401813596317,
	435: -4.888658618255058,
	436: -3.14e+100,
	437: -3.14e+100,
	438: -2.7635336784127853,
	439: -10.275268591948773,
	440: -3.14e+100,
	441: -3.14e+100,
	442: -3.14e+100,
	443: -3.14e+100,
	444: -6.272842531880403,
	445: -6.940320595827818,
	446: -7.728230161053767,
	447: -7.5394037026636855,
	448: -6.85251045118004,
	449: -8.4153713175535,
	450: -8.15808672228609,
	451: -9.299258625372996,
	452: -3.053292303412302,
	453: -3.14e+100,
	454: -5.9430181843676895,
	455: -3.14e+100,
	456: -11.453923588290419,
	457: -3.14e+100,
	458: -3.14e+100,
	459: -8.427419656069674,
	460: -6.1970794699489575,
	461: -13.533365129970255,
	462: -3.14e+100,
	463: -3.14e+100,
}
