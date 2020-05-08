//Struct Model

package main

//import "time"

//--------------------------------------[ struct ]-------------------------------------------------------------------

type RequestThailandPatientInfo struct {
	Data []struct {
		ConfirmDate string `json:"ConfirmDate"`
		No          string `json:"No"`
		Age         int    `json:"Age"`
		Gender      string `json:"Gender"`
		GenderEn    string `json:"GenderEn"`
		Nation      string `json:"Nation"`
		NationEn    string `json:"NationEn"`
		ProvinceId  int    `json:"ProvinceId"`
		Province    string `json:"Province"`
		ProvinceEn  string `json:"ProvinceEn"`
		District    string `json:"District"`
	}
}

type ThailandPatientInfo struct {
	//Id            *string `gorm:"column:id; primary_key"`
	ConfirmDate   *string `gorm:"column:confirm_date;"`
	Age           *int    `gorm:"column:age;"`
	GenderTh      *string `gorm:"column:gender_th;"`
	GenderEn      *string `gorm:"column:gender_en;"`
	NationalityTh *string `gorm:"column:nationality_th;"`
	NationalityEn *string `gorm:"column:nationality_en;"`
	District      *string `gorm:"column:district;"`
	ProvinceId    *int    `gorm:"column:province_id;"`
}

func (ThailandPatientInfo) TableName() string {
	return "THAILAND_PATIENTS_INFO"
}

type RequestTotalThailandPatients struct {
	Confirmed       *int64  `json:"Confirmed"`
	Hospitalized    *int64  `json:"Hospitalized"`
	Recovered       *int64  `json:"Recovered"`
	Deaths          *int64  `json:"Deaths"`
	NewConfirmed    *int64  `json:"NewConfirmed"`
	NewHospitalized *int64  `json:"NewHospitalized"`
	NewRecovered    *int64  `json:"NewRecovered"`
	NewDeaths       *int64  `json:"NewDeaths"`
	UpdateDate      *string `json:"UpdateDate"`
}

type RespondTotalGlobalPatients struct {
	CountryId                 int    `gorm:"column:ct_id; primary_key"`
	TotalCases                int64  `gorm:"column:total_cases;"`
	TotalActiveCases          int64  `gorm:"column:total_active_cases;"`
	TotalRecovered            int64  `gorm:"column:total_recovered;"`
	TotalDeaths               int64  `gorm:"column:total_deaths;"`
	TotalCasesIncreases       int64  `gorm:"column:total_cases_increases;"`
	TotalActiveCasesIncreases int64  `gorm:"column:Total_active_cases_increases;"`
	TotalRecoveredIncreases   int64  `gorm:"column:total_recovered_increases;"`
	TotalDeathsIncreases      int64  `gorm:"column:total_deaths_increases;"`
	UpdateDate                string `gorm:"column:update_date;"`
}

func (RespondTotalGlobalPatients) TableName() string {
	return "TOTAL_GLOBAL_PATIENTS"
}

type UpdateTotalGlobalPatients struct {
	//CountryId				  int  `gorm:"column:ct_id;"` //เอาออกด้วย for Create
	TotalCases                *int64  `gorm:"column:total_cases;"`
	TotalActiveCases          *int64  `gorm:"column:total_active_cases;"`
	TotalRecovered            *int64  `gorm:"column:total_recovered;"`
	TotalDeaths               *int64  `gorm:"column:total_deaths;"`
	TotalCasesIncreases       *int64  `gorm:"column:total_cases_increases;"`
	TotalActiveCasesIncreases *int64  `gorm:"column:Total_active_cases_increases;"`
	TotalRecoveredIncreases   *int64  `gorm:"column:total_recovered_increases;"`
	TotalDeathsIncreases      *int64  `gorm:"column:total_deaths_increases;"`
	UpdateDate                *string `gorm:"column:update_date;"`
}

type Country struct {
	Id          int    `gorm:"column:id;"`
	ctCode		string `gorm:"column:ct_code;"`
	Code        string `gorm:"column:code;"`
	CountryTh   string `gorm:"column:ct_nameTH;"`
	CountryEng  string `gorm:"column:ct_nameEN;"`
	Emoji       string `gorm:"column:emoji;"`
}

func (Country) TableName() string {
	return "COUNTRY"
}

type RequestTotalGlobalPatients struct {
	Countryitems []struct {
		Num1   TotalCountryPatients `json:"1"`
		Num2   TotalCountryPatients `json:"2"`
		Num3   TotalCountryPatients `json:"3"`
		Num4   TotalCountryPatients `json:"4"`
		Num5   TotalCountryPatients `json:"5"`
		Num6   TotalCountryPatients `json:"6"`
		Num7   TotalCountryPatients `json:"7"`
		Num8   TotalCountryPatients `json:"8"`
		Num9   TotalCountryPatients `json:"9"`
		Num10  TotalCountryPatients `json:"10"`
		Num11  TotalCountryPatients `json:"11"`
		Num12  TotalCountryPatients `json:"12"`
		Num13  TotalCountryPatients `json:"13"`
		Num14  TotalCountryPatients `json:"14"`
		Num15  TotalCountryPatients `json:"15"`
		Num16  TotalCountryPatients `json:"16"`
		Num17  TotalCountryPatients `json:"17"`
		Num18  TotalCountryPatients `json:"18"`
		Num19  TotalCountryPatients `json:"19"`
		Num20  TotalCountryPatients `json:"20"`
		Num21  TotalCountryPatients `json:"21"`
		Num22  TotalCountryPatients `json:"22"`
		Num23  TotalCountryPatients `json:"23"`
		Num24  TotalCountryPatients `json:"24"`
		Num25  TotalCountryPatients `json:"25"`
		Num26  TotalCountryPatients `json:"26"`
		Num27  TotalCountryPatients `json:"27"`
		Num28  TotalCountryPatients `json:"28"`
		Num29  TotalCountryPatients `json:"29"`
		Num30  TotalCountryPatients `json:"30"`
		Num31  TotalCountryPatients `json:"31"`
		Num32  TotalCountryPatients `json:"32"`
		Num33  TotalCountryPatients `json:"33"`
		Num34  TotalCountryPatients `json:"34"`
		Num35  TotalCountryPatients `json:"35"`
		Num36  TotalCountryPatients `json:"36"`
		Num37  TotalCountryPatients `json:"37"`
		Num38  TotalCountryPatients `json:"38"`
		Num39  TotalCountryPatients `json:"39"`
		Num40  TotalCountryPatients `json:"40"`
		Num41  TotalCountryPatients `json:"41"`
		Num42  TotalCountryPatients `json:"42"`
		Num43  TotalCountryPatients `json:"43"`
		Num44  TotalCountryPatients `json:"44"`
		Num45  TotalCountryPatients `json:"45"`
		Num46  TotalCountryPatients `json:"46"`
		Num47  TotalCountryPatients `json:"47"`
		Num48  TotalCountryPatients `json:"48"`
		Num49  TotalCountryPatients `json:"49"`
		Num50  TotalCountryPatients `json:"50"`
		Num51  TotalCountryPatients `json:"51"`
		Num52  TotalCountryPatients `json:"52"`
		Num53  TotalCountryPatients `json:"53"`
		Num54  TotalCountryPatients `json:"54"`
		Num55  TotalCountryPatients `json:"55"`
		Num56  TotalCountryPatients `json:"56"`
		Num57  TotalCountryPatients `json:"57"`
		Num58  TotalCountryPatients `json:"58"`
		Num59  TotalCountryPatients `json:"59"`
		Num60  TotalCountryPatients `json:"60"`
		Num61  TotalCountryPatients `json:"61"`
		Num62  TotalCountryPatients `json:"62"`
		Num63  TotalCountryPatients `json:"63"`
		Num64  TotalCountryPatients `json:"64"`
		Num65  TotalCountryPatients `json:"65"`
		Num66  TotalCountryPatients `json:"66"`
		Num67  TotalCountryPatients `json:"67"`
		Num68  TotalCountryPatients `json:"68"`
		Num69  TotalCountryPatients `json:"69"`
		Num70  TotalCountryPatients `json:"70"`
		Num71  TotalCountryPatients `json:"71"`
		Num72  TotalCountryPatients `json:"72"`
		Num73  TotalCountryPatients `json:"73"`
		Num74  TotalCountryPatients `json:"74"`
		Num75  TotalCountryPatients `json:"75"`
		Num76  TotalCountryPatients `json:"76"`
		Num77  TotalCountryPatients `json:"77"`
		Num78  TotalCountryPatients `json:"78"`
		Num79  TotalCountryPatients `json:"79"`
		Num80  TotalCountryPatients `json:"80"`
		Num81  TotalCountryPatients `json:"81"`
		Num82  TotalCountryPatients `json:"82"`
		Num83  TotalCountryPatients `json:"83"`
		Num84  TotalCountryPatients `json:"84"`
		Num85  TotalCountryPatients `json:"85"`
		Num86  TotalCountryPatients `json:"86"`
		Num87  TotalCountryPatients `json:"87"`
		Num88  TotalCountryPatients `json:"88"`
		Num89  TotalCountryPatients `json:"89"`
		Num90  TotalCountryPatients `json:"90"`
		Num91  TotalCountryPatients `json:"91"`
		Num92  TotalCountryPatients `json:"92"`
		Num93  TotalCountryPatients `json:"93"`
		Num94  TotalCountryPatients `json:"94"`
		Num95  TotalCountryPatients `json:"95"`
		Num96  TotalCountryPatients `json:"96"`
		Num97  TotalCountryPatients `json:"97"`
		Num98  TotalCountryPatients `json:"98"`
		Num99  TotalCountryPatients `json:"99"`
		Num100 TotalCountryPatients `json:"100"`
		Num101 TotalCountryPatients `json:"101"`
		Num102 TotalCountryPatients `json:"102"`
		Num103 TotalCountryPatients `json:"103"`
		Num104 TotalCountryPatients `json:"104"`
		Num105 TotalCountryPatients `json:"105"`
		Num106 TotalCountryPatients `json:"106"`
		Num107 TotalCountryPatients `json:"107"`
		Num108 TotalCountryPatients `json:"108"`
		Num109 TotalCountryPatients `json:"109"`
		Num110 TotalCountryPatients `json:"110"`
		Num111 TotalCountryPatients `json:"111"`
		Num112 TotalCountryPatients `json:"112"`
		Num113 TotalCountryPatients `json:"113"`
		Num114 TotalCountryPatients `json:"114"`
		Num115 TotalCountryPatients `json:"115"`
		Num116 TotalCountryPatients `json:"116"`
		Num117 TotalCountryPatients `json:"117"`
		Num118 TotalCountryPatients `json:"118"`
		Num119 TotalCountryPatients `json:"119"`
		Num120 TotalCountryPatients `json:"120"`
		Num121 TotalCountryPatients `json:"121"`
		Num122 TotalCountryPatients `json:"122"`
		Num123 TotalCountryPatients `json:"123"`
		Num124 TotalCountryPatients `json:"124"`
		Num125 TotalCountryPatients `json:"125"`
		Num126 TotalCountryPatients `json:"126"`
		Num127 TotalCountryPatients `json:"127"`
		Num128 TotalCountryPatients `json:"128"`
		Num129 TotalCountryPatients `json:"129"`
		Num130 TotalCountryPatients `json:"130"`
		Num131 TotalCountryPatients `json:"131"`
		Num132 TotalCountryPatients `json:"132"`
		Num133 TotalCountryPatients `json:"133"`
		Num134 TotalCountryPatients `json:"134"`
		Num135 TotalCountryPatients `json:"135"`
		Num136 TotalCountryPatients `json:"136"`
		Num137 TotalCountryPatients `json:"137"`
		Num138 TotalCountryPatients `json:"138"`
		Num139 TotalCountryPatients `json:"139"`
		Num140 TotalCountryPatients `json:"140"`
		Num141 TotalCountryPatients `json:"141"`
		Num142 TotalCountryPatients `json:"142"`
		Num143 TotalCountryPatients `json:"143"`
		Num144 TotalCountryPatients `json:"144"`
		Num145 TotalCountryPatients `json:"145"`
		Num146 TotalCountryPatients `json:"146"`
		Num147 TotalCountryPatients `json:"147"`
		Num148 TotalCountryPatients `json:"148"`
		Num149 TotalCountryPatients `json:"149"`
		Num150 TotalCountryPatients `json:"150"`
		Num151 TotalCountryPatients `json:"151"`
		Num152 TotalCountryPatients `json:"152"`
		Num153 TotalCountryPatients `json:"153"`
		Num154 TotalCountryPatients `json:"154"`
		Num155 TotalCountryPatients `json:"155"`
		Num156 TotalCountryPatients `json:"156"`
		Num157 TotalCountryPatients `json:"157"`
		Num158 TotalCountryPatients `json:"158"`
		Num159 TotalCountryPatients `json:"159"`
		Num160 TotalCountryPatients `json:"160"`
		Num161 TotalCountryPatients `json:"161"`
		Num162 TotalCountryPatients `json:"162"`
		Num163 TotalCountryPatients `json:"163"`
		Num164 TotalCountryPatients `json:"164"`
		Num165 TotalCountryPatients `json:"165"`
		Num166 TotalCountryPatients `json:"166"`
		Num167 TotalCountryPatients `json:"167"`
		Num168 TotalCountryPatients `json:"168"`
		Num169 TotalCountryPatients `json:"169"`
		Num170 TotalCountryPatients `json:"170"`
		Num171 TotalCountryPatients `json:"171"`
		Num172 TotalCountryPatients `json:"172"`
		Num173 TotalCountryPatients `json:"173"`
		Num174 TotalCountryPatients `json:"174"`
		Num175 TotalCountryPatients `json:"175"`
		Num176 TotalCountryPatients `json:"176"`
		Num177 TotalCountryPatients `json:"177"`
		Num178 TotalCountryPatients `json:"178"`
		Num179 TotalCountryPatients `json:"179"`
		Num180 TotalCountryPatients `json:"180"`
		Num181 TotalCountryPatients `json:"181"`
		Num182 TotalCountryPatients `json:"182"`
	} `json:"Countryitems"`
}

type TotalCountryPatients struct {
	Code                string `json:"code"`
	Ourid               int    `json:"ourid"`
	Source              string `json:"source"`
	Title               string `json:"title"`
	TotalActiveCases    int    `json:"total_active_cases"`
	TotalCases          int    `json:"total_cases"`
	TotalDeaths         int    `json:"total_deaths"`
	TotalNewCasesToday  int    `json:"total_new_cases_today"`
	TotalNewDeathsToday int    `json:"total_new_deaths_today"`
	TotalRecovered      int    `json:"total_recovered"`
	TotalSeriousCases   int    `json:"total_serious_cases"`
	TotalUnresolved     int    `json:"total_unresolved"`
}

type RequestTotalThailandPatientsProvince struct {
	Province struct {
		Bangkok               int64 `json:"Bangkok"`
		Unknown               int64 `json:"Unknown"`
		Nonthaburi            int64 `json:"Nonthaburi"`
		Phuket                int64 `json:"Phuket"`
		SamutPrakan           int64 `json:"Samut Prakan"`
		Chonburi              int64 `json:"Chonburi"`
		Yala                  int64 `json:"Yala"`
		Pattani               int64 `json:"Pattani"`
		ChiangMai             int64 `json:"Chiang Mai"`
		Songkhla              int64 `json:"Songkhla"`
		PathumThani           int64 `json:"Pathum Thani"`
		NakhonPathom          int64 `json:"Nakhon Pathom"`
		SamutSakhon           int64 `json:"Samut Sakhon"`
		Chachoengsao          int64 `json:"Chachoengsao"`
		PrachuapKhiriKhan     int64 `json:"Prachuap Khiri Khan"`
		Narathiwat            int64 `json:"Narathiwat"`
		SuratThani            int64 `json:"Surat Thani"`
		NakhonRatchasima      int64 `json:"Nakhon Ratchasima"`
		UbonRatchathani       int64 `json:"Ubon Ratchathani"`
		Krabi                 int64 `json:"Krabi"`
		Buriram               int64 `json:"Buriram"`
		SaKaeo                int64 `json:"Sa Kaeo"`
		Kanchanaburi          int64 `json:"Kanchanaburi"`
		ChiangRai             int64 `json:"Chiang Rai"`
		UdonThani             int64 `json:"Udon Thani"`
		Ratchaburi            int64 `json:"Ratchaburi"`
		Surin                 int64 `json:"Surin"`
		Sisaket               int64 `json:"Sisaket"`
		NakhonSawan           int64 `json:"Nakhon Sawan"`
		Trang                 int64 `json:"Trang"`
		NakhonSiThammarat     int64 `json:"Nakhon Si Thammarat"`
		Phatthalung           int64 `json:"Phatthalung"`
		Rayong                int64 `json:"Rayong"`
		Saraburi              int64 `json:"Saraburi"`
		MaeHongSon            int64 `json:"Mae Hong Son"`
		Tak                   int64 `json:"Tak"`
		Phitsanulok           int64 `json:"Phitsanulok"`
		NongBuaLamphu         int64 `json:"Nong Bua Lamphu"`
		KhonKaen              int64 `json:"Khon Kaen"`
		SuphanBuri            int64 `json:"Suphan Buri"`
		Prachinburi           int64 `json:"Prachinburi"`
		Mukdahan              int64 `json:"Mukdahan"`
		Lamphun               int64 `json:"Lamphun"`
		Phetchabun            int64 `json:"Phetchabun"`
		Sukhothai             int64 `json:"Sukhothai"`
		Chanthaburi           int64 `json:"Chanthaburi"`
		Kalasin               int64 `json:"Kalasin"`
		Uttaradit             int64 `json:"Uttaradit"`
		RoiEt                 int64 `json:"Roi Et"`
		NakhonNayok           int64 `json:"Nakhon Nayok"`
		Phayao                int64 `json:"Phayao"`
		Chaiyaphum            int64 `json:"Chaiyaphum"`
		NongKhai              int64 `json:"Nong Khai"`
		Lopburi               int64 `json:"Lopburi"`
		Loei                  int64 `json:"Loei"`
		AmnatCharoen          int64 `json:"Amnat Charoen"`
		Phetchaburi           int64 `json:"Phetchaburi"`
		Yasothon              int64 `json:"Yasothon"`
		PhraNakhonSiAyutthaya int64 `json:"Phra Nakhon Si Ayutthaya"`
		MahaSarakham          int64 `json:"Maha Sarakham"`
		Phrae                 int64 `json:"Phrae"`
		UthaiThani            int64 `json:"Uthai Thani"`
		Chumphon              int64 `json:"Chumphon"`
		SamutSongkhram        int64 `json:"Samut Songkhram"`
		NakhonPhanom          int64 `json:"Nakhon Phanom"`
		Lampang               int64 `json:"Lampang"`

		//ยังไม่มีผู้ป่วย
		AngThong      int64 `json:"Ang Thong"`
		SingBuri      int64 `json:"Sing Buri"`
		Chainat       int64 `json:"Chainat"`
		Trat          int64 `json:"Trat"`
		SakonNakhon   int64 `json:"Sakon Nakhon"`
		Nan           int64 `json:"Nan"`
		KamphaengPhet int64 `json:"Kamphaeng Phet"`
		Phichit       int64 `json:"Phichit"`
		PhangNga      int64 `json:"Phang Nga"`
		Ranong        int64 `json:"Ranong"`
		Satun         int64 `json:"Satun"`
		BuengKan      int64 `json:"Bueng Kan"`
	}

}

type TotalThailandPatientsProvince struct {
	TotalCase  *int64  `gorm:"column:total_case;"`
}

func (TotalThailandPatientsProvince) TableName() string {
	return "TOTAL_THAILAND_PATIENTS_PROVINCE"
}

type Province struct {
	Id         *int    `gorm:"column:id;"`
	ProvinceTh *string `gorm:"column:province_th;"`
	ProvinceEn *string `gorm:"column:province_en;"`
}

func (Province) TableName() string {
	return "PROVINCE"
}
