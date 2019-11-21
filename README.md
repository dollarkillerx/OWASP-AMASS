# OWASP-AMASS
OWASP-AMASS 源码分析

### 目录分析
整个项目的Go代码大约25039行 并不是太大 就 一个模块一个模块的看
``` 
.
├── alterations
│   ├── alterations.go
│   └── markov.go
├── cmd      程序 入口 通过os.args[1]判断分给不同的模块
│   └── amass
│       ├── amass
│       ├── db.go
│       ├── enum.go
│       ├── intel.go
│       ├── main.go
│       ├── track.go
│       └── viz.go
├── config
│   ├── config.go
│   ├── config-packr.go
│   ├── config_test.go
│   ├── packrd
│   │   └── packed-packr.go
│   └── support.go
├── CONTRIBUTING.md
├── doc
│   ├── install.md
│   └── user_guide.md
├── Dockerfile
├── enum
│   ├── addr.go
│   ├── brute.go
│   ├── enum.go
│   ├── io.go
│   └── names.go
├── eventbus
│   └── bus.go
├── examples
│   ├── config.ini
│   ├── minimal.go
│   └── wordlists
│       ├── all.txt
│       ├── bitquark_subdomains_top100K.txt
│       ├── deepmagic.com_top500prefixes.txt
│       ├── deepmagic.com_top50kprefixes.txt
│       ├── fierce_hostlist.txt
│       ├── jhaddix_all.txt
│       ├── sorted_knock_dnsrecon_fierce_recon-ng.txt
│       ├── subdomains.lst
│       ├── subdomains-top1mil-110000.txt
│       ├── subdomains-top1mil-20000.txt
│       └── subdomains-top1mil-5000.txt
├── format
│   ├── parse.go
│   └── print.go
├── go.mod
├── go.sum
├── graph
│   ├── addr.go
│   ├── as.go
│   ├── db
│   │   ├── cayley.go
│   │   ├── databases.go
│   │   └── gremlin.go
│   ├── event.go
│   ├── fqdn.go
│   ├── graph.go
│   ├── io.go
│   ├── netblock.go
│   ├── source.go
│   └── viz.go
├── images
│   ├── amass.gif
│   ├── maltego_graph_import_wizard.png
│   ├── maltego_results.png
│   ├── network_06062018.png
│   ├── network_06092018.png
│   ├── owasp_logo.png
│   └── snapcraft_icon.png
├── intel
│   └── intel.go
├── LICENSE
├── limits
│   ├── limits_unix.go
│   └── limits_windows.go
├── net
│   ├── dns
│   │   ├── dns.go
│   │   └── dns_test.go
│   ├── http
│   │   └── http.go
│   ├── network.go
│   └── network_test.go
├── queue
│   └── queue.go
├── README.md
├── requests
│   └── request.go
├── resolvers
│   ├── assess.go
│   ├── msgs.go
│   ├── pool.go
│   ├── pool_test.go
│   ├── ratemon.go
│   ├── resolver.go
│   ├── resolver_test.go
│   ├── scored.go
│   ├── wildcards.go
│   └── zone.go
├── resources
│   ├── alterations.txt
│   ├── asnlist.txt
│   ├── namelist.txt
│   ├── tldlist.txt
│   └── user_agents.txt
├── semaphore
│   └── semaphore.go
├── services
│   ├── alienvault.go
│   ├── alienvault_test.go
│   ├── archiveit.go
│   ├── archiveit_test.go
│   ├── archivetoday.go
│   ├── archivetoday_test.go
│   ├── arquivo.go
│   ├── arquivo_test.go
│   ├── ask.go
│   ├── ask_test.go
│   ├── baidu.go
│   ├── baidu_test.go
│   ├── binaryedge.go
│   ├── binaryedge_test.go
│   ├── bing.go
│   ├── bing_test.go
│   ├── bufferover.go
│   ├── bufferover_test.go
│   ├── censys.go
│   ├── censys_test.go
│   ├── certspotter.go
│   ├── certspotter_test.go
│   ├── circl.go
│   ├── circl_test.go
│   ├── commoncrawl.go
│   ├── commoncrawl_test.go
│   ├── crtsh.go
│   ├── crtsh_test.go
│   ├── datamgmtsrv.go
│   ├── dnsdb.go
│   ├── dnsdumpster.go
│   ├── dnsdumpster_test.go
│   ├── dnssrv.go
│   ├── dnstable.go
│   ├── dnstable_test.go
│   ├── dogpile.go
│   ├── dogpile_test.go
│   ├── entrust.go
│   ├── entrust_test.go
│   ├── exalead.go
│   ├── exalead_test.go
│   ├── github.go
│   ├── googlect.go
│   ├── google.go
│   ├── google_test.go
│   ├── hackerone.go
│   ├── hackerone_test.go
│   ├── hackertarget.go
│   ├── hackertarget_test.go
│   ├── ipapi.go
│   ├── ipv4info.go
│   ├── ipv4info_test.go
│   ├── local.go
│   ├── locarchive.go
│   ├── locarchive_test.go
│   ├── mnemonic.go
│   ├── netcraft.go
│   ├── netcraft_test.go
│   ├── networksdb.go
│   ├── openukarchive.go
│   ├── openukarchive_test.go
│   ├── passivetotal.go
│   ├── passivetotal_test.go
│   ├── pastebin.go
│   ├── pastebin_test.go
│   ├── ptrarchive.go
│   ├── ptrarchive_test.go
│   ├── radb.go
│   ├── riddler.go
│   ├── riddler_test.go
│   ├── robtex.go
│   ├── robtex_test.go
│   ├── securitytrails.go
│   ├── securitytrails_test.go
│   ├── service.go
│   ├── shadowserver.go
│   ├── shodan.go
│   ├── shodan_test.go
│   ├── sitedossier.go
│   ├── sitedossier_test.go
│   ├── sources.go
│   ├── sources_test.go
│   ├── spyse.go
│   ├── spyse_test.go
│   ├── srv_records.go
│   ├── sublist3r.go
│   ├── sublist3r_test.go
│   ├── system.go
│   ├── teamcymru.go
│   ├── threatcrowd.go
│   ├── threatcrowd_test.go
│   ├── twitter.go
│   ├── twitter_test.go
│   ├── ukgovarchive.go
│   ├── ukgovarchive_test.go
│   ├── umbrella.go
│   ├── urlscan.go
│   ├── urlscan_test.go
│   ├── viewdns.go
│   ├── viewdns_test.go
│   ├── virustotal.go
│   ├── virustotal_test.go
│   ├── wayback.go
│   ├── wayback_test.go
│   ├── whoisxml.go
│   ├── yahoo.go
│   └── yahoo_test.go
├── stringset
│   ├── filter.go
│   ├── filter_test.go
│   ├── set.go
│   └── set_test.go
├── viz
│   ├── d3.go
│   ├── dot.go
│   ├── gexf.go
│   ├── graphistry.go
│   ├── maltego.go
│   ├── visjs.go
│   └── viz.go
└── wordlist
    ├── wordlist.go
    └── wordlist_test.go
```

### 要分析一个项目 先要了解这个是干什么的 并灵活的使用  一来就看源码 她怎么用的你都不知道
- 基础命令

主命令的入口在:
`cmd/main.go - 86`
``` bash
amass enum -d example.com
```

这里对用户参数进行分流  到不同的模块下
`main 101` 
``` go
	// 获取用户输入参数进行匹配开头   // 将余下的给到第二层   (例如 ./maass enum  [-help  这里是余下的给到了第二层])
	switch os.Args[1] {
	case "db":        // 图形数据库
		runDBCommand(os.Args[2:])
	case "enum":      // 子域名枚举
		runEnumCommand(os.Args[2:])
	case "intel":     // 更具公开资料查询
		runIntelCommand(os.Args[2:])
	case "track":     
		runTrackCommand(os.Args[2:])
	case "viz":       // 枚举可视化
		runVizCommand(os.Args[2:])
	default:
		commandUsage(mainUsageMsg, mainFlagSet, defaultBuf)
		os.Exit(1)
	}
```

- cmd 下的 enum 模块 主要职责 负责enum所需参数的初始化和validate
  完成初始化后 交付给enum包下的enum进一步处理

### amass的進程守護用了sighup來作 
