package data

// Company consists of company information
var Company = map[string][]string{
	"name":      {"3 Round Stones, Inc.", "48 Factoring Inc.", "5PSolutions", "Abt Associates", "Accela", "Accenture", "AccuWeather", "Acxiom", "Adaptive", "Adobe Digital Government", "Aidin", "Alarm.com", "Allianz", "Allied Van Lines", "AllState Insurance Group", "Alltuition", "Altova", "Amazon Web Services", "American Red Ball Movers", "Amida Technology Solutions", "Analytica", "Apextech LLC", "Appallicious", "Aquicore", "Archimedes Inc.", "AreaVibes Inc.", "Arpin Van Lines", "Arrive Labs", "ASC Partners", "Asset4", "Atlas Van Lines", "AtSite", "Aunt Bertha, Inc.", "Aureus Sciences (*Now part of Elsevier)", "AutoGrid Systems", "Avalara", "Avvo", "Ayasdi", "Azavea", "BaleFire Global", "Barchart", "Be Informed", "Bekins", "Berkery Noyes MandASoft", "Berkshire Hathaway", "BetterLesson", "BillGuard", "Bing", "Biovia", "BizVizz", "BlackRock", "Bloomberg", "Booz Allen Hamilton", "Boston Consulting Group", "Boundless", "Bridgewater", "Brightscope", "BuildFax", "Buildingeye", "BuildZoom", "Business and Legal Resources", "Business Monitor International", "Calcbench, Inc.", "Cambridge Information Group", "Cambridge Semantics", "CAN Capital", "Canon", "Capital Cube", "Cappex", "Captricity", "CareSet Systems", "Careset.com", "CARFAX", "Caspio", "Castle Biosciences", "CB Insights", "Ceiba Solutions", "Center for Responsive Politics", "Cerner", "Certara", "CGI", "Charles River Associates", "Charles Schwab Corp.", "Chemical Abstracts Service", "Child Care Desk", "Chubb", "Citigroup", "CityScan", "CitySourced", "Civic Impulse LLC", "Civic Insight", "Civinomics", "Civis Analytics", "Clean Power Finance", "ClearHealthCosts", "ClearStory Data", "Climate Corporation", "CliniCast", "Cloudmade", "Cloudspyre", "Code for America", "Code-N", "Collective IP", "College Abacus, an ECMC initiative", "College Board", "Compared Care", "Compendia Bioscience Life Technologies", "Compliance and Risks", "Computer Packages Inc", "CONNECT-DOT LLC.", "ConnectEDU", "Connotate", "Construction Monitor LLC", "Consumer Reports", "CoolClimate", "Copyright Clearance Center", "CoreLogic", "CostQuest", "Credit Karma", "Credit Sesame", "CrowdANALYTIX", "Dabo Health", "DataLogix", "DataMade", "DataMarket", "Datamyne", "DataWeave", "Deloitte", "DemystData", "Department of Better Technology", "Development Seed", "Docket Alarm, Inc.", "Dow Jones & Co.", "Dun & Bradstreet", "Earth Networks", "EarthObserver App", "Earthquake Alert!", "Eat Shop Sleep", "Ecodesk", "eInstitutional", "Embark", "EMC", "Energy Points, Inc.", "Energy Solutions Forum", "Enervee Corporation", "Enigma.io", "Ensco", "Environmental Data Resources", "Epsilon", "Equal Pay for Women", "Equifax", "Equilar", "Ernst & Young LLP", "eScholar LLC.", "Esri", "Estately", "Everyday Health", "Evidera", "Experian", "Expert Health Data Programming, Inc.", "Exversion", "Ez-XBRL", "Factset", "Factual", "Farmers", "FarmLogs", "Fastcase", "Fidelity Investments", "FindTheBest.com", "First Fuel Software", "FirstPoint, Inc.", "Fitch", "FlightAware", "FlightStats", "FlightView", "Food+Tech Connect", "Forrester Research", "Foursquare", "Fujitsu", "Funding Circle", "FutureAdvisor", "Fuzion Apps, Inc.", "Gallup", "Galorath Incorporated", "Garmin", "Genability", "GenoSpace", "Geofeedia", "Geolytics", "Geoscape", "GetRaised", "GitHub", "Glassy Media", "Golden Helix", "GoodGuide", "Google Maps", "Google Public Data Explorer", "Government Transaction Services", "Govini", "GovTribe", "Govzilla, Inc.", "gRadiant Research LLC", "Graebel Van Lines", "Graematter, Inc.", "Granicus", "GreatSchools", "GuideStar", "H3 Biomedicine", "Harris Corporation", "HDScores, Inc", "Headlight", "Healthgrades", "Healthline", "HealthMap", "HealthPocket, Inc.", "HelloWallet", "HERE", "Honest Buildings", "HopStop", "Housefax", "How's My Offer?", "IBM", "ideas42", "iFactor Consulting", "IFI CLAIMS Patent Services", "iMedicare", "Impact Forecasting (Aon)", "Impaq International", "Import.io", "IMS Health", "InCadence", "indoo.rs", "InfoCommerce Group", "Informatica", "InnoCentive", "Innography", "Innovest Systems", "Inovalon", "Inrix Traffic", "Intelius", "Intermap Technologies", "Investormill", "Iodine", "IPHIX", "iRecycle", "iTriage", "IVES Group Inc", "IW Financial", "JJ Keller", "J.P. Morgan Chase", "Junar, Inc.", "Junyo", "Jurispect", "Kaiser Permanante", "karmadata", "Keychain Logistics Corp.", "KidAdmit, Inc.", "Kimono Labs", "KLD Research", "Knoema", "Knowledge Agency", "KPMG", "Kroll Bond Ratings Agency", "Kyruus", "Lawdragon", "Legal Science Partners", "(Leg)Cyte", "LegiNation, Inc.", "LegiStorm", "Lenddo", "Lending Club", "Level One Technologies", "LexisNexis", "Liberty Mutual Insurance Cos.", "Lilly Open Innovation Drug Discovery", "Liquid Robotics", "Locavore", "LOGIXDATA, LLC", "LoopNet", "Loqate, Inc.", "LoseIt.com", "LOVELAND Technologies", "Lucid", "Lumesis, Inc.", "Mango Transit", "Mapbox", "Maponics", "MapQuest", "Marinexplore, Inc.", "MarketSense", "Marlin & Associates", "Marlin Alter and Associates", "McGraw Hill Financial", "McKinsey", "MedWatcher", "Mercaris", "Merrill Corp.", "Merrill Lynch", "MetLife", "mHealthCoach", "MicroBilt Corporation", "Microsoft Windows Azure Marketplace", "Mint", "Moody's", "Morgan Stanley", "Morningstar, Inc.", "Mozio", "MuckRock.com", "Munetrix", "Municode", "National Van Lines", "Nationwide Mutual Insurance Company", "Nautilytics", "Navico", "NERA Economic Consulting", "NerdWallet", "New Media Parents", "Next Step Living", "NextBus", "nGAP Incorporated", "Nielsen", "Noesis", "NonprofitMetrics", "North American Van Lines", "Noveda Technologies", "NuCivic", "Numedii", "Oliver Wyman", "OnDeck", "OnStar", "Ontodia, Inc", "Onvia", "Open Data Nation", "OpenCounter", "OpenGov", "OpenPlans", "OpportunitySpace, Inc.", "Optensity", "optiGov", "OptumInsight", "Orlin Research", "OSIsoft", "OTC Markets", "Outline", "Oversight Systems", "Overture Technologies", "Owler", "Palantir Technologies", "Panjiva", "Parsons Brinckerhoff", "Patently-O", "PatientsLikeMe", "Pave", "Paxata", "PayScale, Inc.", "PeerJ", "People Power", "Persint", "Personal Democracy Media", "Personal, Inc.", "Personalis", "Peterson's", "PEV4me.com", "PIXIA Corp", "PlaceILive.com", "PlanetEcosystems", "PlotWatt", "Plus-U", "PolicyMap", "Politify", "Poncho App", "POPVOX", "Porch", "PossibilityU", "PowerAdvocate", "Practice Fusion", "Predilytics", "PricewaterhouseCoopers (PWC)", "ProgrammableWeb", "Progressive Insurance Group", "Propeller Health", "ProPublica", "PublicEngines", "PYA Analytics", "Qado Energy, Inc.", "Quandl", "Quertle", "Quid", "R R Donnelley", "RAND Corporation", "Rand McNally", "Rank and Filed", "Ranku", "Rapid Cycle Solutions", "realtor.com", "Recargo", "ReciPal", "Redfin", "RedLaser", "Reed Elsevier", "REI Systems", "Relationship Science", "Remi", "Retroficiency", "Revaluate", "Revelstone", "Rezolve Group", "Rivet Software", "Roadify Transit", "Robinson + Yu", "Russell Investments", "Sage Bionetworks", "SAP", "SAS", "Scale Unlimited", "Science Exchange", "Seabourne", "SeeClickFix", "SigFig", "Simple Energy", "SimpleTuition", "SlashDB", "Smart Utility Systems", "SmartAsset", "SmartProcure", "Smartronix", "SnapSense", "Social Explorer", "Social Health Insights", "SocialEffort Inc", "Socrata", "Solar Census", "SolarList", "Sophic Systems Alliance", "S&P Capital IQ", "SpaceCurve", "SpeSo Health", "Spikes Cavell Analytic Inc", "Splunk", "Spokeo", "SpotCrime", "SpotHero.com", "Stamen Design", "Standard and Poor's", "State Farm Insurance", "Sterling Infosystems", "Stevens Worldwide Van Lines", "STILLWATER SUPERCOMPUTING INC", "StockSmart", "Stormpulse", "StreamLink Software", "StreetCred Software, Inc", "StreetEasy", "Suddath", "Symcat", "Synthicity", "T. Rowe Price", "Tableau Software", "TagniFi", "Telenav", "Tendril", "Teradata", "The Advisory Board Company", "The Bridgespan Group", "The DocGraph Journal", "The Govtech Fund", "The Schork Report", "The Vanguard Group", "Think Computer Corporation", "Thinknum", "Thomson Reuters", "TopCoder", "TowerData", "TransparaGov", "TransUnion", "TrialTrove", "TrialX", "Trintech", "TrueCar", "Trulia", "TrustedID", "TuvaLabs", "Uber", "Unigo LLC", "United Mayflower", "Urban Airship", "Urban Mapping, Inc", "US Green Data", "U.S. News Schools", "USAA Group", "USSearch", "Verdafero", "Vimo", "VisualDoD, LLC", "Vital Axiom | Niinja", "VitalChek", "Vitals", "Vizzuality", "Votizen", "Walk Score", "WaterSmart Software", "WattzOn", "Way Better Patents", "Weather Channel", "Weather Decision Technologies", "Weather Underground", "WebFilings", "Webitects", "WebMD", "Weight Watchers", "WeMakeItSafer", "Wheaton World Wide Moving", "Whitby Group", "Wolfram Research", "Wolters Kluwer", "Workhands", "Xatori", "Xcential", "xDayta", "Xignite", "Yahoo", "Zebu Compliance Solutions", "Yelp", "YourMapper", "Zillow", "ZocDoc", "Zonability", "Zoner", "Zurich Insurance (Risk Room)"},
	"suffix":    {"Inc", "and Sons", "LLC", "Group"},
	"buzzwords": {"Adaptive", "Advanced", "Ameliorated", "Assimilated", "Automated", "Balanced", "Business-focused", "Centralized", "Cloned", "Compatible", "Configurable", "Cross-group", "Cross-platform", "Customer-focused", "Customizable", "De-engineered", "Decentralized", "Devolved", "Digitized", "Distributed", "Diverse", "Down-sized", "Enhanced", "Enterprise-wide", "Ergonomic", "Exclusive", "Expanded", "Extended", "Face to face", "Focused", "Front-line", "Fully-configurable", "Function-based", "Fundamental", "Future-proofed", "Grass-roots", "Horizontal", "Implemented", "Innovative", "Integrated", "Intuitive", "Inverse", "Managed", "Mandatory", "Monitored", "Multi-channelled", "Multi-lateral", "Multi-layered", "Multi-tiered", "Networked", "Object-based", "Open-architected", "Open-source", "Operative", "Optimized", "Optional", "Organic", "Organized", "Persevering", "Persistent", "Phased", "Polarised", "Pre-emptive", "Proactive", "Profit-focused", "Profound", "Programmable", "Progressive", "Public-key", "Quality-focused", "Re-contextualized", "Re-engineered", "Reactive", "Realigned", "Reduced", "Reverse-engineered", "Right-sized", "Robust", "Seamless", "Secured", "Self-enabling", "Sharable", "Stand-alone", "Streamlined", "Switchable", "Synchronised", "Synergistic", "Synergized", "Team-oriented", "Total", "Triple-buffered", "Universal", "Up-sized", "Upgradable", "User-centric", "User-friendly", "Versatile", "Virtual", "Vision-oriented", "Visionary", "24 hour", "24/7", "3rd generation", "4th generation", "5th generation", "6th generation", "actuating", "analyzing", "asymmetric", "asynchronous", "attitude-oriented", "background", "bandwidth-monitored", "bi-directional", "bifurcated", "bottom-line", "clear-thinking", "client-driven", "client-server", "coherent", "cohesive", "composite", "content-based", "context-sensitive", "contextually-based", "dedicated", "demand-driven", "didactic", "directional", "discrete", "disintermediate", "dynamic", "eco-centric", "empowering", "encompassing", "even-keeled", "executive", "explicit", "exuding", "fault-tolerant", "foreground", "fresh-thinking", "full-range", "global", "grid-enabled", "heuristic", "high-level", "holistic", "homogeneous", "human-resource", "hybrid", "impactful", "incremental", "intangible", "interactive", "intermediate", "leading edge", "local", "logistical", "maximized", "methodical", "mission-critical", "mobile", "modular", "motivating", "multi-state", "multi-tasking", "multimedia", "national", "needs-based", "neutral", "next generation", "non-volatile", "object-oriented", "optimal", "optimizing", "radical", "real-time", "reciprocal", "regional", "responsive", "scalable", "secondary", "solution-oriented", "stable", "static", "system-worthy", "systematic", "systemic", "tangible", "tertiary", "transitional", "uniform", "upward-trending", "user-facing", "value-added", "web-enabled", "well-modulated", "zero administration", "zero defect", "zero tolerance", "Graphic Interface", "Graphical User Interface", "ability", "access", "adapter", "algorithm", "alliance", "analyzer", "application", "approach", "architecture", "archive", "array", "artificial intelligence", "attitude", "benchmark", "budgetary management", "capability", "capacity", "challenge", "circuit", "collaboration", "complexity", "concept", "conglomeration", "contingency", "core", "customer loyalty", "data-warehouse", "database", "definition", "emulation", "encoding", "encryption", "extranet", "firmware", "flexibility", "focus group", "forecast", "frame", "framework", "function", "functionalities", "groupware", "hardware", "help-desk", "hierarchy", "hub", "implementation", "info-mediaries", "infrastructure", "initiative", "installation", "instruction set", "interface", "internet solution", "intranet", "knowledge base", "knowledge user", "leverage", "local area network", "matrices", "matrix", "methodology", "middleware", "migration", "model", "moderator", "monitoring", "moratorium", "neural-net", "open architecture", "open system", "orchestration", "paradigm", "parallelism", "policy", "portal", "pricing structure", "process improvement", "product", "productivity", "project", "projection", "protocol", "secured line", "service-desk", "software", "solution", "standardization", "strategy", "structure", "success", "superstructure", "support", "synergy", "system engine", "task-force", "throughput", "time-frame", "toolset", "utilisation", "website", "workforce"},
	"bs":        {"aggregate", "architect", "benchmark", "brand", "cultivate", "deliver", "deploy", "disintermediate", "drive", "e-enable", "embrace", "empower", "enable", "engage", "engineer", "enhance", "envisioneer", "evolve", "expedite", "exploit", "extend", "facilitate", "generate", "grow", "harness", "implement", "incentivize", "incubate", "innovate", "integrate", "iterate", "leverage", "matrix", "maximize", "mesh", "monetize", "morph", "optimize", "orchestrate", "productize", "recontextualize", "redefine", "reintermediate", "reinvent", "repurpose", "revolutionize", "scale", "seize", "strategize", "streamline", "syndicate", "synergize", "synthesize", "target", "transform", "transition", "unleash", "utilize", "visualize", "whiteboard", "24-365", "24-7", "B2B", "B2C", "back-end", "best-of-breed", "bleeding-edge", "bricks-and-clicks", "clicks-and-mortar", "collaborative", "compelling", "cross-media", "cross-platform", "customized", "cutting-edge", "distributed", "dot-com", "dynamic", "e-business", "efficient", "end-to-end", "enterprise", "extensible", "frictionless", "front-end", "global", "granular", "holistic", "impactful", "innovative", "integrated", "interactive", "intuitive", "killer", "leading-edge", "magnetic", "mission-critical", "next-generation", "one-to-one", "open-source", "out-of-the-box", "plug-and-play", "proactive", "real-time", "revolutionary", "rich", "robust", "scalable", "seamless", "sexy", "sticky", "strategic", "synergistic", "transparent", "turn-key", "ubiquitous", "user-centric", "value-added", "vertical", "viral", "virtual", "visionary", "web-enabled", "wireless", "world-class", "ROI", "action-items", "applications", "architectures", "bandwidth", "channels", "communities", "content", "convergence", "deliverables", "e-business", "e-commerce", "e-markets", "e-services", "e-tailers", "experiences", "eyeballs", "functionalities", "infomediaries", "infrastructures", "initiatives", "interfaces", "markets", "methodologies", "metrics", "mindshare", "models", "networks", "niches", "paradigms", "partnerships", "platforms", "portals", "relationships", "schemas", "solutions", "supply-chains", "synergies", "systems", "technologies", "users", "vortals", "web-services", "web-readiness"},
}
