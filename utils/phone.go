package utils

import (
	"regexp"
	"strconv"
	"strings"
)

var dialCode = map[string]struct {
	Code   string
	Length int
}{
	"Afghanistan":                      {"+93", 9},
	"Albania":                          {"+355", 9},
	"Algeria":                          {"+213", 9},
	"American Samoa":                   {"+1-684", 7},
	"Andorra":                          {"+376", 9},
	"Angola":                           {"+244", 9},
	"Anguilla":                         {"+1-264", 7},
	"Antarctica":                       {"+672", 6},
	"Antigua and Barbuda":              {"+1-268", 7},
	"Argentina":                        {"+54", 9},
	"Armenia":                          {"+374", 8},
	"Aruba":                            {"+297", 7},
	"Australia":                        {"+61", 9},
	"Austria":                          {"+43", 10},
	"Azerbaijan":                       {"+994", 9},
	"Bahamas":                          {"+1-242", 7},
	"Bahrain":                          {"+973", 8},
	"Bangladesh":                       {"+880", 10},
	"Barbados":                         {"+1-246", 7},
	"Belarus":                          {"+375", 9},
	"Belgium":                          {"+32", 9},
	"Belize":                           {"+501", 7},
	"Benin":                            {"+229", 8},
	"Bermuda":                          {"+1-441", 7},
	"Bhutan":                           {"+975", 8},
	"Bolivia":                          {"+591", 8},
	"Bosnia and Herzegovina":           {"+387", 8},
	"Botswana":                         {"+267", 8},
	"Brazil":                           {"+55", 11},
	"British Indian Ocean Territory":   {"+246", 7},
	"British Virgin Islands":           {"+1-284", 7},
	"Brunei":                           {"+673", 8},
	"Bulgaria":                         {"+359", 9},
	"Burkina Faso":                     {"+226", 8},
	"Burundi":                          {"+257", 8},
	"Cambodia":                         {"+855", 9},
	"Cameroon":                         {"+237", 9},
	"Canada":                           {"+1", 10},
	"Cape Verde":                       {"+238", 7},
	"Cayman Islands":                   {"+1-345", 7},
	"Central African Republic":         {"+236", 8},
	"Chad":                             {"+235", 8},
	"Chile":                            {"+56", 9},
	"China":                            {"+86", 11},
	"Christmas Island":                 {"+61", 9},
	"Cocos Islands":                    {"+61", 9},
	"Colombia":                         {"+57", 10},
	"Comoros":                          {"+269", 8},
	"Cook Islands":                     {"+682", 5},
	"Costa Rica":                       {"+506", 8},
	"Croatia":                          {"+385", 9},
	"Cuba":                             {"+53", 8},
	"Curacao":                          {"+599", 8},
	"Cyprus":                           {"+357", 9},
	"Czech Republic":                   {"+420", 9},
	"Democratic Republic of the Congo": {"+243", 9},
	"Denmark":                          {"+45", 8},
	"Djibouti":                         {"+253", 8},
	"Dominica":                         {"+1-767", 7},
	// "Dominican Republic":               {"+1-809, +1-829, +1-849", 10},
	"East Timor":               {"+670", 8},
	"Ecuador":                  {"+593", 9},
	"Egypt":                    {"+20", 9},
	"El Salvador":              {"+503", 8},
	"Equatorial Guinea":        {"+240", 9},
	"Eritrea":                  {"+291", 8},
	"Estonia":                  {"+372", 9},
	"Ethiopia":                 {"+251", 9},
	"Falkland Islands":         {"+500", 7},
	"Faroe Islands":            {"+298", 6},
	"Fiji":                     {"+679", 7},
	"Finland":                  {"+358", 10},
	"France":                   {"+33", 9},
	"French Polynesia":         {"+689", 7},
	"Gabon":                    {"+241", 8},
	"Gambia":                   {"+220", 7},
	"Georgia":                  {"+995", 9},
	"Germany":                  {"+49", 11},
	"Ghana":                    {"+233", 10},
	"Gibraltar":                {"+350", 8},
	"Greece":                   {"+30", 10},
	"Greenland":                {"+299", 6},
	"Grenada":                  {"+1-473", 7},
	"Guam":                     {"+1-671", 7},
	"Guatemala":                {"+502", 8},
	"Guernsey":                 {"+44-1481", 9},
	"Guinea":                   {"+224", 8},
	"Guinea-Bissau":            {"+245", 7},
	"Guyana":                   {"+592", 7},
	"Haiti":                    {"+509", 8},
	"Honduras":                 {"+504", 8},
	"Hong Kong":                {"+852", 8},
	"Hungary":                  {"+36", 9},
	"Iceland":                  {"+354", 7},
	"India":                    {"+91", 10},
	"Indonesia":                {"+62", 11},
	"Iran":                     {"+98", 10},
	"Iraq":                     {"+964", 10},
	"Ireland":                  {"+353", 9},
	"Isle of Man":              {"+44-1624", 10},
	"Israel":                   {"+972", 9},
	"Italy":                    {"+39", 10},
	"Ivory Coast":              {"+225", 8},
	"Jamaica":                  {"+1-876", 7},
	"Japan":                    {"+81", 11},
	"Jersey":                   {"+44-1534", 9},
	"Jordan":                   {"+962", 9},
	"Kazakhstan":               {"+7", 10},
	"Kenya":                    {"+254", 9},
	"Kiribati":                 {"+686", 5},
	"Kosovo":                   {"+383", 8},
	"Kuwait":                   {"+965", 8},
	"Kyrgyzstan":               {"+996", 9},
	"Laos":                     {"+856", 9},
	"Latvia":                   {"+371", 8},
	"Lebanon":                  {"+961", 8},
	"Lesotho":                  {"+266", 8},
	"Liberia":                  {"+231", 7},
	"Libya":                    {"+218", 9},
	"Liechtenstein":            {"+423", 9},
	"Lithuania":                {"+370", 8},
	"Luxembourg":               {"+352", 9},
	"Macau":                    {"+853", 8},
	"Macedonia":                {"+389", 8},
	"Madagascar":               {"+261", 9},
	"Malawi":                   {"+265", 9},
	"Malaysia":                 {"+60", 10},
	"Maldives":                 {"+960", 7},
	"Mali":                     {"+223", 8},
	"Malta":                    {"+356", 8},
	"Marshall Islands":         {"+692", 7},
	"Mauritania":               {"+222", 8},
	"Mauritius":                {"+230", 7},
	"Mayotte":                  {"+262", 9},
	"Mexico":                   {"+52", 10},
	"Micronesia":               {"+691", 7},
	"Moldova":                  {"+373", 8},
	"Monaco":                   {"+377", 9},
	"Mongolia":                 {"+976", 8},
	"Montenegro":               {"+382", 8},
	"Montserrat":               {"+1-664", 7},
	"Morocco":                  {"+212", 9},
	"Mozambique":               {"+258", 9},
	"Myanmar":                  {"+95", 9},
	"Namibia":                  {"+264", 9},
	"Nauru":                    {"+674", 7},
	"Nepal":                    {"+977", 10},
	"Netherlands":              {"+31", 10},
	"Netherlands Antilles":     {"+599", 8},
	"New Caledonia":            {"+687", 7},
	"New Zealand":              {"+64", 10},
	"Nicaragua":                {"+505", 8},
	"Niger":                    {"+227", 8},
	"Nigeria":                  {"+234", 10},
	"Niue":                     {"+683", 4},
	"North Korea":              {"+850", 8},
	"Northern Mariana Islands": {"+1-670", 7},
	"Norway":                   {"+47", 8},
	"Oman":                     {"+968", 8},
	"Pakistan":                 {"+92", 10},
	"Palau":                    {"+680", 7},
	"Palestine":                {"+970", 9},
	"Panama":                   {"+507", 8},
	"Papua New Guinea":         {"+675", 8},
	"Paraguay":                 {"+595", 9},
	"Peru":                     {"+51", 9},
	"Philippines":              {"+63", 10},
	"Pitcairn":                 {"+64", 4},
	"Poland":                   {"+48", 9},
	"Portugal":                 {"+351", 9},
	// "Puerto Rico":                      {"+1-787, +1-939", 10},
	"Qatar":                            {"+974", 8},
	"Republic of the Congo":            {"+242", 9},
	"Reunion":                          {"+262", 9},
	"Romania":                          {"+40", 9},
	"Russia":                           {"+7", 11},
	"Rwanda":                           {"+250", 9},
	"Saint Barthelemy":                 {"+590", 10},
	"Saint Helena":                     {"+290", 5},
	"Saint Kitts and Nevis":            {"+1-869", 7},
	"Saint Lucia":                      {"+1-758", 7},
	"Saint Martin":                     {"+590", 10},
	"Saint Pierre and Miquelon":        {"+508", 7},
	"Saint Vincent and the Grenadines": {"+1-784", 7},
	"Samoa":                            {"+685", 7},
	"San Marino":                       {"+378", 7},
	"Sao Tome and Principe":            {"+239", 7},
	"Saudi Arabia":                     {"+966", 9},
	"Senegal":                          {"+221", 9},
	"Serbia":                           {"+381", 10},
	"Seychelles":                       {"+248", 7},
	"Sierra Leone":                     {"+232", 8},
	"Singapore":                        {"+65", 8},
	"Sint Maarten":                     {"+1-721", 7},
	"Slovakia":                         {"+421", 9},
	"Slovenia":                         {"+386", 8},
	"Solomon Islands":                  {"+677", 7},
	"Somalia":                          {"+252", 8},
	"South Africa":                     {"+27", 9},
	"South Korea":                      {"+82", 10},
	"South Sudan":                      {"+211", 9},
	"Spain":                            {"+34", 9},
	"Sri Lanka":                        {"+94", 10},
	"Sudan":                            {"+249", 9},
	"Suriname":                         {"+597", 7},
	"Svalbard and Jan Mayen":           {"+47", 8},
	"Swaziland":                        {"+268", 8},
	"Sweden":                           {"+46", 10},
	"Switzerland":                      {"+41", 9},
	"Syria":                            {"+963", 10},
	"Taiwan":                           {"+886", 10},
	"Tajikistan":                       {"+992", 9},
	"Tanzania":                         {"+255", 10},
	"Thailand":                         {"+66", 10},
	"Togo":                             {"+228", 8},
	"Tokelau":                          {"+690", 4},
	"Tonga":                            {"+676", 5},
	"Trinidad and Tobago":              {"+1-868", 7},
	"Tunisia":                          {"+216", 8},
	"Turkey":                           {"+90", 10},
	"Turkmenistan":                     {"+993", 8},
	"Turks and Caicos Islands":         {"+1-649", 7},
	"Tuvalu":                           {"+688", 5},
	"U.S. Virgin Islands":              {"+1-340", 7},
	"Uganda":                           {"+256", 9},
	"Ukraine":                          {"+380", 9},
	"United Arab Emirates":             {"+971", 9},
	"United Kingdom":                   {"+44", 11},
	"United States":                    {"+1", 10},
	"Uruguay":                          {"+598", 8},
	"Uzbekistan":                       {"+998", 9},
	"Vanuatu":                          {"+678", 7},
	"Vatican":                          {"+379", 7},
	"Venezuela":                        {"+58", 10},
	"Vietnam":                          {"+84", 10},
	"Wallis and Futuna":                {"+681", 6},
	"Western Sahara":                   {"+212", 9},
	"Yemen":                            {"+967", 9},
	"Zambia":                           {"+260", 9},
	"Zimbabwe":                         {"+263", 9},
}

func ValidatePhone(dCode []string, val string) string {
	regex := regexp.MustCompile("\\s+")
	phoneVal := regex.ReplaceAllString(val, "")

	if len(dCode) > 0 {
		l, isNumber := strconv.Atoi(dCode[0])
		if isNumber != nil {
			// If the user has not provided a number(length of input),then the user must have provided a country name
			// Check if country exists
			_, ok := dialCode[dCode[0]]
			if !ok {
				return "Invalid country code"
			}

			// Check if phone number is valid
			if len(phoneVal) != (dialCode[dCode[0]].Length+len(dialCode[dCode[0]].Code)) && !strings.HasPrefix(val, dialCode[dCode[0]].Code) {
				return "Invalid phone number"
			}
		}

		var maxLen int
		var err error

		// Then there must be a min length and a max length
		if len(dCode) == 2 {
			maxLen, err = strconv.Atoi(dCode[1])
			if err != nil {
				return "Invalid max length"
			}
		}

		// Simply ignore the case if max length provided is less than the min length
		if maxLen < l {
			maxLen = l
		}

		if len(phoneVal) < l || len(phoneVal) > maxLen {
			return "Invalid phone number"
		}
	}

	// pattern := `\d`
	re := regexp.MustCompile("\\d")
	if !re.MatchString(phoneVal) {
		return "Invalid phone number"
	}

	return ""
}
