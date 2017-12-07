package main

import (
	"fmt"

	"adventofcode2017/advent2017"
)

var (
	input01 = "5994521226795838486188872189952551475352929145357284983463678944777228139398117649129843853837124228353689551178129353548331779783742915361343229141538334688254819714813664439268791978215553677772838853328835345484711229767477729948473391228776486456686265114875686536926498634495695692252159373971631543594656954494117149294648876661157534851938933954787612146436571183144494679952452325989212481219139686138139314915852774628718443532415524776642877131763359413822986619312862889689472397776968662148753187767793762654133429349515324333877787925465541588584988827136676376128887819161672467142579261995482731878979284573246533688835226352691122169847832943513758924194232345988726741789247379184319782387757613138742817826316376233443521857881678228694863681971445442663251423184177628977899963919997529468354953548612966699526718649132789922584524556697715133163376463256225181833257692821331665532681288216949451276844419154245423434141834913951854551253339785533395949815115622811565999252555234944554473912359674379862182425695187593452363724591541992766651311175217218144998691121856882973825162368564156726989939993412963536831593196997676992942673571336164535927371229823236937293782396318237879715612956317715187757397815346635454412183198642637577528632393813964514681344162814122588795865169788121655353319233798811796765852443424783552419541481132132344487835757888468196543736833342945718867855493422435511348343711311624399744482832385998592864795271972577548584967433917322296752992127719964453376414665576196829945664941856493768794911984537445227285657716317974649417586528395488789946689914972732288276665356179889783557481819454699354317555417691494844812852232551189751386484638428296871436139489616192954267794441256929783839652519285835238736142997245189363849356454645663151314124885661919451447628964996797247781196891787171648169427894282768776275689124191811751135567692313571663637214298625367655969575699851121381872872875774999172839521617845847358966264291175387374464425566514426499166813392768677233356646752273398541814142523651415521363267414564886379863699323887278761615927993953372779567675"
	input02 = input01
	input03 = `4347	3350	196	162	233	4932	4419	3485	4509	4287	4433	4033	207	3682	2193	4223
648	94	778	957	1634	2885	1964	2929	2754	89	972	112	80	2819	543	2820
400	133	1010	918	1154	1008	126	150	1118	117	148	463	141	940	1101	89
596	527	224	382	511	565	284	121	643	139	625	335	657	134	125	152
2069	1183	233	213	2192	193	2222	2130	2073	2262	1969	2159	2149	410	181	1924
1610	128	1021	511	740	1384	459	224	183	266	152	1845	1423	230	1500	1381
5454	3936	250	5125	244	720	5059	202	4877	5186	313	6125	172	727	1982	748
3390	3440	220	228	195	4525	1759	1865	1483	5174	4897	4511	5663	4976	3850	199
130	1733	238	1123	231	1347	241	291	1389	1392	269	1687	1359	1694	1629	1750
1590	1394	101	434	1196	623	1033	78	890	1413	74	1274	1512	1043	1103	84
203	236	3001	1664	195	4616	2466	4875	986	1681	152	3788	541	4447	4063	5366
216	4134	255	235	1894	5454	1529	4735	4962	220	2011	2475	185	5060	4676	4089
224	253	19	546	1134	3666	3532	207	210	3947	2290	3573	3808	1494	4308	4372
134	130	2236	118	142	2350	3007	2495	2813	2833	2576	2704	169	2666	2267	850
401	151	309	961	124	1027	1084	389	1150	166	1057	137	932	669	590	188
784	232	363	316	336	666	711	430	192	867	628	57	222	575	622	234`
	input04 = input03
)

func main() {
	fmt.Println(advent2017.SumSkipDigits(input01, 1))
	fmt.Println(advent2017.SumSkipDigits(input02, len(input02)/2))
	fmt.Println(advent2017.RowMaxDiffChecksum(input03))
	fmt.Println(advent2017.RowDivChecksum(input04))
}
