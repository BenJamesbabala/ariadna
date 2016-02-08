package importer

var ESSettings string = `{
    "settings": {
        "analysis": {
            "filter": {
                "map_poi_filter": {
                    "type": "synonym",
                    "synonyms": [
                    	"минфин,министертсво финансов",
                    	"минздрав,министерство здравоохранения",
                    	"минюст,министерство юстиции",
                    	"минтранспорта,министерство транспорта",
                    	"минобразования,министерство образования",
                    	"минкультуры,министерство культуры",
                    	"юракадемия,юридическая академия",
                    	"к-т,кинотеатр",
                    	"маг,магазин",
                    	"тц,торговый центр",
                    	"трк, торгово-развлекательный центр",
                    	"маг,магазин",
						"мц,медицинский центр",
						"мед центр,медицинский центр",
						"мкр,микрорайон",
						"нац (госпиталь),национальный (госпиталь)",
						"дет сад,детский сад",
						"ж/м,жилмассив",
						"жм,жилмассив",
						"КТР",
						"НБТ",
						"к/р,кинотеатр",
						"КНУ",
						"БГУ",
						"г-ца,гостиница",
						"гост,гостиница",
						"МВД,министерство внутренних дел",
						"АУЦА, Американский университет в центральной азии",
						"РОВД,районный отдел внутренних дел",
						"первом,первомайский",
						"сверд,свердловский",
						"октябрь,октябрьский",
						"ГУВД,государственное управление внутренних дел",
						"спец больница,специализированная больница",
						"с,село",
						"кисб,КИКБ",
						"бц,бизнес центр",
						"жд,железнодорожный",
						"сто,станция технического обслуживания",
						"азс,автомобильная заправочная станция",
						"заправка, автомобильная заправочная станция",
						"НХЦ,национальный хирургический центр",
						"ШВК,ШампанВинКомбинат",
						"шампанкомбинат,ШампанВинКомбинат",
						"ГПТ,ГлавПивТрест",
						"БЧК,большой чуйский канал",
						"ДК,дом культуры",
						"МИД,министерство иностранных дел",
						"роддом,родильный дом",
						"гор,городская",
						"корп,корпорация",
						"гор гаи,гаи города бишкек",
						"мед академия,кыргызская государственная медицинская академия",
						"юр академия,кыргызская государственная юридическая академия",
						"воен городок,военный городок",
						"кож завод,кожевенный завод",
						"вечерка,редакция газеты вечерний бишкек",
						"л толстого,льва толстого",
						"д асановой,динары асановой",
						"цсм,центр семейной медицины",
						"центр,центральный",
						"аламед,аламединский",
						"универ,университет",
						"ун-т,университет",
						"гм,гипермаркет",
						"дет больница,детская больница",
						"главпочтамт,главное почтовое отделение",
						"пер,переулок",
						"д сяопина,дэн сяопина",
						"мин юст,министерство юстиции",
						"морфо корпус,морг",
						"энергосбыт,Бишкекский Энергосбыт, ОсОО СеверЭлектро",
						"образ центр,образовательный центральный",
						"гсин,государственная служба исполнения наказаний",
						"респуб,республиканский",
						"обл,областной, областная",
						"мчс,министерство чрезвычайных ситуаций",
						"с батора,сухэ батора",
						"нац банк,национальный банк кыргызской республики",
						"нбкр,национальный банк кыргызской республики",
						"юр фак,юридический факультет",
						"гум фак,гуманитарный факультет",
						"эконом фак,экономический факультет",
						"ген прокуратура,генеральная прокуратура",
						"бгтс,бишкекская городская телефонная сеть",
						"к кийская,кызыл кийская",
						"пив бар,пивной бар",
						"юж ворота,южные ворота"
                    ]
                }
            },
            "analyzer": {
                "map_synonyms": {
                    "tokenizer": "standard",
                    "filter": [
                        "lowercase",
                        "map_poi_filter"
                    ]
                }
            }
        }
    },
    "mappings":
        {"address":
            {"properties":
                {
                    "centroid": {
                        "type": "geo_point"
                    }

                }
            }
        }
}`
