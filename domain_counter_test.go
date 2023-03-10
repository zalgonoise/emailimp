package customerimporter_test

import (
	"encoding/csv"
	"errors"
	"os"
	"testing"

	. "github.com/zalgonoise/emailimp"
)

const (
	rawPath           = "./testdata/customers.csv"
	invalidDomainPath = "./testdata/invalid_domain.csv"
	invalidCSVPath    = "./testdata/invalid_csv.csv"
	invalidPath       = "/no/way/this/dir/exists_"
)

func TestParse(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		entries, err := Parse(rawPath)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(entries) != len(expectedResults) {
			t.Errorf("output length mismatch error: wanted %d ; got %d", len(expectedResults), len(entries))
		}

		for _, e := range entries {
			if expectedResults[e.Domain] != e.Count {
				t.Errorf("output mismatch error: expected domain %s to have %d users ; has %d", e.Domain, expectedResults[e.Domain], e.Count)
			}
		}
	})

	t.Run("Fail", func(t *testing.T) {
		t.Run("InvalidDomain", func(t *testing.T) {
			_, err := Parse(invalidDomainPath)
			if err == nil {
				t.Error("expected an error; got nil")
				return
			}
			if !errors.Is(err, ErrInvalidDomain) {
				t.Errorf("unexpected error: wanted %v ; got %v", ErrInvalidDomain, err)
				return
			}
		})

		t.Run("InvalidCSV", func(t *testing.T) {
			_, err := Parse(invalidCSVPath)
			if err == nil {
				t.Error("expected an error; got nil")
				return
			}
			if !errors.Is(err, csv.ErrFieldCount) {
				t.Errorf("unexpected error: wanted %v ; got %v", csv.ErrFieldCount, err)
				return
			}
		})

		t.Run("InvalidPath", func(t *testing.T) {
			_, err := Parse(invalidPath)
			if err == nil {
				t.Error("expected an error; got nil")
				return
			}
			if !errors.Is(err, os.ErrNotExist) {
				t.Errorf("unexpected error: wanted %v ; got %v", ErrInvalidDomain, err)
				return
			}
		})
	})
}

var expectedResults = map[string]int{
	"123-reg.co.uk":          8,
	"163.com":                6,
	"1688.com":               3,
	"1und1.de":               5,
	"360.cn":                 6,
	"4shared.com":            5,
	"51.la":                  4,
	"a8.net":                 6,
	"abc.net.au":             7,
	"about.com":              5,
	"about.me":               2,
	"aboutads.info":          2,
	"accuweather.com":        6,
	"acquirethisname.com":    6,
	"addthis.com":            10,
	"addtoany.com":           6,
	"admin.ch":               9,
	"adobe.com":              6,
	"alexa.com":              6,
	"alibaba.com":            7,
	"altervista.org":         7,
	"amazon.co.jp":           5,
	"amazon.co.uk":           3,
	"amazon.com":             6,
	"amazon.de":              7,
	"amazonaws.com":          6,
	"ameblo.jp":              6,
	"angelfire.com":          5,
	"answers.com":            5,
	"aol.com":                5,
	"apache.org":             2,
	"apple.com":              5,
	"archive.org":            7,
	"arizona.edu":            5,
	"army.mil":               8,
	"arstechnica.com":        9,
	"artisteer.com":          5,
	"ask.com":                6,
	"auda.org.au":            5,
	"baidu.com":              7,
	"bandcamp.com":           3,
	"barnesandnoble.com":     8,
	"bbb.org":                5,
	"bbc.co.uk":              10,
	"behance.net":            10,
	"berkeley.edu":           7,
	"biblegateway.com":       6,
	"bigcartel.com":          2,
	"biglobe.ne.jp":          4,
	"bing.com":               4,
	"bizjournals.com":        2,
	"blinklist.com":          7,
	"blog.com":               6,
	"blogger.com":            8,
	"bloglines.com":          6,
	"bloglovin.com":          7,
	"blogs.com":              4,
	"blogspot.com":           6,
	"blogtalkradio.com":      7,
	"bloomberg.com":          9,
	"bluehost.com":           5,
	"booking.com":            4,
	"boston.com":             4,
	"bravesites.com":         5,
	"businessinsider.com":    5,
	"businessweek.com":       3,
	"businesswire.com":       2,
	"buzzfeed.com":           10,
	"ca.gov":                 9,
	"cafepress.com":          6,
	"cam.ac.uk":              5,
	"canalblog.com":          7,
	"cargocollective.com":    6,
	"cbc.ca":                 7,
	"cbslocal.com":           8,
	"cbsnews.com":            8,
	"cdbaby.com":             4,
	"cdc.gov":                6,
	"census.gov":             7,
	"chicagotribune.com":     7,
	"china.com.cn":           5,
	"chron.com":              6,
	"chronoengine.com":       7,
	"cisco.com":              11,
	"clickbank.net":          5,
	"cloudflare.com":         3,
	"cmu.edu":                4,
	"cnbc.com":               6,
	"cnet.com":               6,
	"cnn.com":                4,
	"cocolog-nifty.com":      7,
	"columbia.edu":           4,
	"com.com":                5,
	"comcast.net":            3,
	"comsenz.com":            7,
	"constantcontact.com":    8,
	"cornell.edu":            4,
	"cpanel.net":             4,
	"craigslist.org":         8,
	"creativecommons.org":    5,
	"csmonitor.com":          6,
	"cyberchimps.com":        6,
	"dagondesign.com":        10,
	"dailymail.co.uk":        9,
	"dailymotion.com":        1,
	"de.vu":                  4,
	"dedecms.com":            4,
	"delicious.com":          6,
	"deliciousdays.com":      7,
	"dell.com":               4,
	"desdev.cn":              11,
	"devhub.com":             6,
	"deviantart.com":         7,
	"digg.com":               2,
	"diigo.com":              5,
	"dion.ne.jp":             7,
	"discovery.com":          5,
	"discuz.net":             5,
	"disqus.com":             6,
	"dmoz.org":               5,
	"domainmarket.com":       13,
	"dot.gov":                4,
	"dropbox.com":            5,
	"drupal.org":             3,
	"dyndns.org":             4,
	"e-recht24.de":           7,
	"earthlink.net":          9,
	"ebay.co.uk":             5,
	"ebay.com":               4,
	"economist.com":          3,
	"ed.gov":                 3,
	"edublogs.org":           8,
	"eepurl.com":             7,
	"ehow.com":               7,
	"elegantthemes.com":      9,
	"elpais.com":             10,
	"engadget.com":           6,
	"epa.gov":                8,
	"etsy.com":               6,
	"europa.eu":              2,
	"eventbrite.com":         4,
	"examiner.com":           5,
	"example.com":            8,
	"exblog.jp":              5,
	"ezinearticles.com":      9,
	"facebook.com":           7,
	"fastcompany.com":        4,
	"fc2.com":                6,
	"fda.gov":                2,
	"feedburner.com":         5,
	"fema.gov":               5,
	"flavors.me":             7,
	"flickr.com":             5,
	"forbes.com":             7,
	"fotki.com":              5,
	"foxnews.com":            5,
	"free.fr":                8,
	"freewebs.com":           10,
	"friendfeed.com":         4,
	"ft.com":                 5,
	"ftc.gov":                9,
	"furl.net":               4,
	"g.co":                   6,
	"geocities.com":          12,
	"geocities.jp":           6,
	"github.com":             7,
	"github.io":              8,
	"gizmodo.com":            5,
	"globo.com":              10,
	"gmpg.org":               5,
	"gnu.org":                5,
	"go.com":                 11,
	"godaddy.com":            5,
	"goo.gl":                 8,
	"goo.ne.jp":              6,
	"goodreads.com":          3,
	"google.ca":              7,
	"google.cn":              5,
	"google.co.jp":           5,
	"google.co.uk":           10,
	"google.com.au":          8,
	"google.com":             9,
	"google.com.br":          5,
	"google.com.hk":          5,
	"google.de":              11,
	"google.es":              9,
	"google.fr":              12,
	"google.it":              4,
	"google.nl":              7,
	"google.pl":              6,
	"google.ru":              9,
	"gov.uk":                 4,
	"gravatar.com":           4,
	"guardian.co.uk":         8,
	"hao123.com":             7,
	"harvard.edu":            8,
	"hatena.ne.jp":           4,
	"hc360.com":              5,
	"hexun.com":              8,
	"hhs.gov":                5,
	"hibu.com":               10,
	"histats.com":            5,
	"home.pl":                3,
	"homestead.com":          2,
	"hostgator.com":          9,
	"house.gov":              5,
	"howstuffworks.com":      5,
	"hp.com":                 4,
	"hubpages.com":           9,
	"hud.gov":                6,
	"huffingtonpost.com":     10,
	"hugedomains.com":        6,
	"i2i.jp":                 9,
	"ibm.com":                5,
	"icio.us":                3,
	"icq.com":                8,
	"ifeng.com":              9,
	"ihg.com":                9,
	"illinois.edu":           6,
	"imageshack.us":          8,
	"imdb.com":               2,
	"imgur.com":              2,
	"independent.co.uk":      11,
	"indiatimes.com":         5,
	"indiegogo.com":          8,
	"infoseek.co.jp":         4,
	"instagram.com":          6,
	"intel.com":              4,
	"irs.gov":                9,
	"is.gd":                  5,
	"issuu.com":              7,
	"istockphoto.com":        2,
	"jalbum.net":             2,
	"japanpost.jp":           5,
	"java.com":               2,
	"jiathis.com":            10,
	"jigsy.com":              2,
	"jimdo.com":              5,
	"joomla.org":             2,
	"jugem.jp":               5,
	"kickstarter.com":        10,
	"last.fm":                6,
	"latimes.com":            2,
	"linkedin.com":           2,
	"list-manage.com":        3,
	"live.com":               7,
	"liveinternet.ru":        7,
	"livejournal.com":        8,
	"loc.gov":                14,
	"lulu.com":               7,
	"lycos.com":              6,
	"mac.com":                7,
	"macromedia.com":         7,
	"mail.ru":                5,
	"mapquest.com":           5,
	"mapy.cz":                5,
	"marketwatch.com":        7,
	"marriott.com":           6,
	"mashable.com":           5,
	"mayoclinic.com":         5,
	"mediafire.com":          4,
	"meetup.com":             4,
	"merriam-webster.com":    6,
	"microsoft.com":          7,
	"miibeian.gov.cn":        4,
	"miitbeian.gov.cn":       12,
	"mit.edu":                9,
	"mlb.com":                7,
	"moonfruit.com":          6,
	"mozilla.com":            7,
	"mozilla.org":            2,
	"msn.com":                6,
	"msu.edu":                3,
	"mtv.com":                6,
	"multiply.com":           5,
	"myspace.com":            8,
	"mysql.com":              6,
	"narod.ru":               2,
	"nasa.gov":               8,
	"nationalgeographic.com": 6,
	"nature.com":             7,
	"naver.com":              5,
	"nba.com":                5,
	"nbcnews.com":            5,
	"netlog.com":             8,
	"netscape.com":           4,
	"netvibes.com":           8,
	"networkadvertising.org": 11,
	"networksolutions.com":   3,
	"newsvine.com":           8,
	"newyorker.com":          3,
	"nhs.uk":                 9,
	"nifty.com":              9,
	"nih.gov":                6,
	"ning.com":               1,
	"noaa.gov":               6,
	"npr.org":                3,
	"nps.gov":                3,
	"nsw.gov.au":             6,
	"nydailynews.com":        7,
	"nymag.com":              11,
	"nytimes.com":            3,
	"nyu.edu":                4,
	"oaic.gov.au":            6,
	"oakley.com":             9,
	"ocn.ne.jp":              4,
	"odnoklassniki.ru":       6,
	"omniture.com":           5,
	"opensource.org":         5,
	"opera.com":              10,
	"oracle.com":             5,
	"over-blog.com":          6,
	"ovh.net":                9,
	"ow.ly":                  6,
	"ox.ac.uk":               3,
	"pagesperso-orange.fr":   4,
	"paginegialle.it":        6,
	"parallels.com":          6,
	"patch.com":              8,
	"paypal.com":             5,
	"pbs.org":                6,
	"pcworld.com":            4,
	"pen.io":                 4,
	"people.com.cn":          4,
	"phoca.cz":               6,
	"photobucket.com":        9,
	"php.net":                2,
	"phpbb.com":              6,
	"pinterest.com":          7,
	"plala.or.jp":            4,
	"posterous.com":          3,
	"princeton.edu":          6,
	"printfriendly.com":      6,
	"privacy.gov.au":         4,
	"prlog.org":              6,
	"prnewswire.com":         7,
	"prweb.com":              8,
	"psu.edu":                6,
	"purevolume.com":         8,
	"qq.com":                 5,
	"quantcast.com":          3,
	"rakuten.co.jp":          8,
	"rambler.ru":             10,
	"redcross.org":           5,
	"reddit.com":             13,
	"rediff.com":             9,
	"reference.com":          4,
	"reuters.com":            6,
	"reverbnation.com":       6,
	"sakura.ne.jp":           7,
	"salon.com":              5,
	"samsung.com":            5,
	"sbwire.com":             6,
	"sciencedaily.com":       5,
	"sciencedirect.com":      4,
	"scientificamerican.com": 4,
	"scribd.com":             5,
	"seattletimes.com":       10,
	"seesaa.net":             5,
	"senate.gov":             5,
	"sfgate.com":             10,
	"shareasale.com":         6,
	"shinystat.com":          10,
	"shop-pro.jp":            3,
	"shutterfly.com":         6,
	"si.edu":                 6,
	"simplemachines.org":     4,
	"sina.com.cn":            9,
	"sitemeter.com":          6,
	"skype.com":              8,
	"skyrock.com":            7,
	"slashdot.org":           8,
	"slate.com":              4,
	"slideshare.net":         9,
	"smh.com.au":             8,
	"smugmug.com":            8,
	"so-net.ne.jp":           7,
	"sogou.com":              7,
	"sohu.com":               4,
	"soundcloud.com":         3,
	"soup.io":                4,
	"sourceforge.net":        7,
	"sphinn.com":             9,
	"spiegel.de":             2,
	"spotify.com":            5,
	"springer.com":           5,
	"squarespace.com":        10,
	"squidoo.com":            9,
	"stanford.edu":           11,
	"statcounter.com":        6,
	"state.gov":              3,
	"state.tx.us":            4,
	"storify.com":            8,
	"studiopress.com":        12,
	"stumbleupon.com":        7,
	"sun.com":                10,
	"surveymonkey.com":       5,
	"symantec.com":           8,
	"t-online.de":            9,
	"t.co":                   3,
	"tamu.edu":               8,
	"taobao.com":             7,
	"techcrunch.com":         4,
	"technorati.com":         3,
	"ted.com":                1,
	"telegraph.co.uk":        7,
	"theatlantic.com":        5,
	"theglobeandmail.com":    5,
	"theguardian.com":        3,
	"themeforest.net":        2,
	"thetimes.co.uk":         4,
	"time.com":               8,
	"timesonline.co.uk":      5,
	"tiny.cc":                9,
	"tinypic.com":            6,
	"tinyurl.com":            9,
	"tmall.com":              6,
	"toplist.cz":             6,
	"topsy.com":              4,
	"trellian.com":           2,
	"tripadvisor.com":        6,
	"tripod.com":             7,
	"tumblr.com":             7,
	"tuttocitta.it":          7,
	"twitpic.com":            8,
	"twitter.com":            8,
	"typepad.com":            8,
	"ucla.edu":               7,
	"ucoz.com":               5,
	"ucoz.ru":                2,
	"ucsd.edu":               8,
	"uiuc.edu":               6,
	"umich.edu":              6,
	"umn.edu":                5,
	"un.org":                 1,
	"unblog.fr":              8,
	"unc.edu":                8,
	"unesco.org":             7,
	"unicef.org":             6,
	"uol.com.br":             5,
	"upenn.edu":              6,
	"usa.gov":                4,
	"usatoday.com":           7,
	"usda.gov":               6,
	"usgs.gov":               5,
	"usnews.com":             8,
	"ustream.tv":             4,
	"utexas.edu":             5,
	"va.gov":                 3,
	"vimeo.com":              3,
	"vinaora.com":            5,
	"virginia.edu":           5,
	"vistaprint.com":         8,
	"vk.com":                 11,
	"vkontakte.ru":           2,
	"w3.org":                 6,
	"walmart.com":            7,
	"washington.edu":         4,
	"washingtonpost.com":     10,
	"weather.com":            5,
	"webeden.co.uk":          4,
	"webmd.com":              5,
	"webnode.com":            5,
	"webs.com":               5,
	"weebly.com":             12,
	"weibo.com":              5,
	"whitehouse.gov":         9,
	"who.int":                3,
	"wikia.com":              6,
	"wikimedia.org":          6,
	"wikipedia.org":          5,
	"wikispaces.com":         4,
	"wiley.com":              8,
	"wired.com":              8,
	"wisc.edu":               8,
	"wix.com":                4,
	"woothemes.com":          10,
	"wordpress.com":          4,
	"wordpress.org":          7,
	"wp.com":                 7,
	"wsj.com":                5,
	"wufoo.com":              2,
	"wunderground.com":       3,
	"xing.com":               9,
	"xinhuanet.com":          6,
	"xrea.com":               2,
	"yahoo.co.jp":            5,
	"yahoo.com":              7,
	"yale.edu":               7,
	"yandex.ru":              5,
	"ycombinator.com":        6,
	"yellowbook.com":         5,
	"yellowpages.com":        7,
	"yelp.com":               4,
	"yolasite.com":           9,
	"youku.com":              5,
	"youtu.be":               6,
	"youtube.com":            3,
	"zdnet.com":              8,
	"zimbio.com":             3,
}
