package batchread

// Hacker consists of random hacker phrases
var Hacker = map[string][]string{
	"abbreviation": {"TCP", "HTTP", "SDD", "RAM", "GB", "CSS", "SSL", "AGP", "SQL", "FTP", "PCI", "AI", "ADP", "RSS", "XML", "EXE", "COM", "HDD", "THX", "SMTP", "SMS", "USB", "PNG", "SAS", "IB", "SCSI", "JSON", "XSS", "JBOD"},
	"adjective":    {"auxiliary", "primary", "back-end", "digital", "open-source", "virtual", "cross-platform", "redundant", "online", "haptic", "multi-byte", "bluetooth", "wireless", "1080p", "neural", "optical", "solid state", "mobile"},
	"noun":         {"driver", "protocol", "bandwidth", "panel", "microchip", "program", "port", "card", "array", "interface", "system", "sensor", "firewall", "hard drive", "pixel", "alarm", "feed", "monitor", "application", "transmitter", "bus", "circuit", "capacitor", "matrix"},
	"verb":         {"back up", "bypass", "hack", "override", "compress", "copy", "navigate", "index", "connect", "generate", "quantify", "calculate", "synthesize", "input", "transmit", "program", "reboot", "parse"},
	"ingverb":      {"backing up", "bypassing", "hacking", "overriding", "compressing", "copying", "navigating", "indexing", "connecting", "generating", "quantifying", "calculating", "synthesizing", "transmitting", "programming", "parsing"},
	"phrase": {
		"If we {hackerverb} the {hackernoun}, we can get to the {hackerabbreviation} {hackernoun} through the {hackeradjective} {hackerabbreviation} {hackernoun}!",
		"We need to {hackerverb} the {hackeradjective} {hackerabbreviation} {hackernoun}!",
		"Try to {hackerverb} the {hackerabbreviation} {hackernoun}, maybe it will {hackerverb} the {hackeradjective} {hackernoun}!",
		"You can't {hackerverb} the {hackernoun} without {hackeringverb} the {hackeradjective} {hackerabbreviation} {hackernoun}!",
		"Use the {hackeradjective} {hackerabbreviation} {hackernoun}, then you can {hackerverb} the {hackeradjective} {hackernoun}!",
		"The {hackerabbreviation} {hackernoun} is down, {hackerverb} the {hackeradjective} {hackernoun} so we can {hackerverb} the {hackerabbreviation} {hackernoun}!",
		"{hackeringverb} the {hackernoun} won't do anything, we need to {hackerverb} the {hackeradjective} {hackerabbreviation} {hackernoun}!",
		"I'll {hackerverb} the {hackeradjective} {hackerabbreviation} {hackernoun}, that should {hackerverb} the {hackerabbreviation} {hackernoun}!",
	},
}

// Beer consists of various beer information
var Beer = map[string][]string{
	"name":  {"Pliny The Elder", "Founders Kentucky Breakfast", "Trappistes Rochefort 10", "HopSlam Ale", "Stone Imperial Russian Stout", "St. Bernardus Abt 12", "Founders Breakfast Stout", "Weihenstephaner Hefeweissbier", "Péché Mortel", "Celebrator Doppelbock", "Duvel", "Dreadnaught IPA", "Nugget Nectar", "La Fin Du Monde", "Bourbon County Stout", "Old Rasputin Russian Imperial Stout", "Two Hearted Ale", "Ruination IPA", "Schneider Aventinus", "Double Bastard Ale", "90 Minute IPA", "Hop Rod Rye", "Trappistes Rochefort 8", "Chimay Grande Réserve", "Stone IPA", "Arrogant Bastard Ale", "Edmund Fitzgerald Porter", "Chocolate St", "Oak Aged Yeti Imperial Stout", "Ten FIDY", "Storm King Stout", "Shakespeare Oatmeal", "Alpha King Pale Ale", "Westmalle Trappist Tripel", "Samuel Smith’s Imperial IPA", "Yeti Imperial Stout", "Hennepin", "Samuel Smith’s Oatmeal Stout", "Brooklyn Black", "Oaked Arrogant Bastard Ale", "Sublimely Self-Righteous Ale", "Trois Pistoles", "Bell’s Expedition", "Sierra Nevada Celebration Ale", "Sierra Nevada Bigfoot Barleywine Style Ale", "Racer 5 India Pale Ale, Bear Republic Bre", "Orval Trappist Ale", "Hercules Double IPA", "Maharaj", "Maudite"},
	"hop":   {"Ahtanum", "Amarillo", "Bitter Gold", "Bravo", "Brewer’s Gold", "Bullion", "Cascade", "Cashmere", "Centennial", "Chelan", "Chinook", "Citra", "Cluster", "Columbia", "Columbus", "Comet", "Crystal", "Equinox", "Eroica", "Fuggle", "Galena", "Glacier", "Golding", "Hallertau", "Horizon", "Liberty", "Magnum", "Millennium", "Mosaic", "Mt. Hood", "Mt. Rainier", "Newport", "Northern Brewer", "Nugget", "Olympic", "Palisade", "Perle", "Saaz", "Santiam", "Simcoe", "Sorachi Ace", "Sterling", "Summit", "Tahoma", "Tettnang", "TriplePearl", "Ultra", "Vanguard", "Warrior", "Willamette", "Yakima Gol"},
	"yeast": {"1007 - German Ale", "1010 - American Wheat", "1028 - London Ale", "1056 - American Ale", "1084 - Irish Ale", "1098 - British Ale", "1099 - Whitbread Ale", "1187 - Ringwood Ale", "1272 - American Ale II", "1275 - Thames Valley Ale", "1318 - London Ale III", "1332 - Northwest Ale", "1335 - British Ale II", "1450 - Dennys Favorite 50", "1469 - West Yorkshire Ale", "1728 - Scottish Ale", "1968 - London ESB Ale", "2565 - Kölsch", "1214 - Belgian Abbey", "1388 - Belgian Strong Ale", "1762 - Belgian Abbey II", "3056 - Bavarian Wheat Blend", "3068 - Weihenstephan Weizen", "3278 - Belgian Lambic Blend", "3333 - German Wheat", "3463 - Forbidden Fruit", "3522 - Belgian Ardennes", "3638 - Bavarian Wheat", "3711 - French Saison", "3724 - Belgian Saison", "3763 - Roeselare Ale Blend", "3787 - Trappist High Gravity", "3942 - Belgian Wheat", "3944 - Belgian Witbier", "2000 - Budvar Lager", "2001 - Urquell Lager", "2007 - Pilsen Lager", "2035 - American Lager", "2042 - Danish Lager", "2112 - California Lager", "2124 - Bohemian Lager", "2206 - Bavarian Lager", "2278 - Czech Pils", "2308 - Munich Lager", "2633 - Octoberfest Lager Blend", "5112 - Brettanomyces bruxellensis", "5335 - Lactobacillus", "5526 - Brettanomyces lambicus", "5733 - Pediococcus"},
	"malt":  {"Black malt", "Caramel", "Carapils", "Chocolate", "Munich", "Caramel", "Carapils", "Chocolate malt", "Munich", "Pale", "Roasted barley", "Rye malt", "Special roast", "Victory", "Vienna", "Wheat mal"},
	"style": {"Light Lager", "Pilsner", "European Amber Lager", "Dark Lager", "Bock", "Light Hybrid Beer", "Amber Hybrid Beer", "English Pale Ale", "Scottish And Irish Ale", "Merican Ale", "English Brown Ale", "Porter", "Stout", "India Pale Ale", "German Wheat And Rye Beer", "Belgian And French Ale", "Sour Ale", "Belgian Strong Ale", "Strong Ale", "Fruit Beer", "Vegetable Beer", "Smoke-flavored", "Wood-aged Beer"},
}

var Data = map[string]map[string][]string{
	"hacker": Hacker,
	"beer":   Beer,
}

// Intn is a part of the randomizer package. It is used to select an item from the database.
type Intn interface {
	Intn(n int) int
}

// GetRandValue from a [category, subcategory]. Returns error if not found.
func GetRandValue(f Intn, category, subcategory string) (string, error) {
	index := f.Intn(len(Data[category][subcategory]))
	return Data[category][subcategory][index], nil
}
