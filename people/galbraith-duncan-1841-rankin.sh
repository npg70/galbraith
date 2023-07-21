name Duncan Galbraith
tags Greenock
external {
    familysearch KZP7-YK7
}
birth -date 26-jun-1841 
baptism -date 14-sep-1846 -location 'Kilcalmonell and Kilberry' -ref duncan1841-baptism
death -date 9-jul-1885 -location greenock -ref duncan1841-death -note "by drowning, there is a correction to death record containing more information."

body {
    He was a police officer.
}

body {
1871 Census of Scotland, Renfrewshire, Greenock,East Shaw Street, 32:$ref[census1871]
$csvtable{
First,Last,Role,Age,Year,Birth
Duncan, Galbreath,Head,28,1843,Kilbery
Mary, Galbreath,Wife,27,1844,Jura
Neil, Galbreath, Son, 2,1879,Greenock
Marion, Galbreath, Dau,0,1871,Greenock
}
}

body {
1881 Census of Scotland, Renfrewshire, Middle, East, Regent Street, 19:$ref[census1881]
$csvtable{
First,Last,Role,Age,Year,Birth
Duncan,Galbraith,Head,36,1845,Kintyre
Mary,Galbraith,Wife,37,1844,Jura
Neil,Galbraith,Son,12,1869,Greenock
Marion S, Galbraith,Dau,10,1871,Greenock
Mary,Galbraith,Dau,8,1873,Greenock
John,Galbraith,Son,4,1877,Greenock
}
}

note {
Marriage record lists his mother as "Mary Sinclair," which doesn't match.  However, Mary Stout's husband name was "Sinclair" so perhaps a mix up.
}

note {
Age at death in 1885 says 41 (or maybe 48?), which would have a birth year of 1844 (or 1837), which does not match baptism records of 1841.
}

partner {
    name Mary Rankin
    marriage -date 16-jan-1868 -location greenock -ref rankin-marriage
    baptism -date 26-jan-1844 -location jura
    death -date 24-feb-1916 -location greenock

    child galbraith-neil-1868-mclean

    child {
        name Marion Shaw Galbraith
        birth -date 1871 -location greenock
        death -date 22-dec-1950 -location greenock -ref marion1871-death
        partner {
            name Alexander Lamont
        }
    }
    child {
        name Mary Galbraith
        birth -date 1873 -location greenock
    }
    child {
        name John Rankin Galbraith
        birth -date 16-jan-1877 -location greenock -ref john1877-birth
        death -date 1881 -location greenock, -ref john1877-death
    }
}

footnotes {

    duncan1841-baptism {
        $opr-ref-link[b-1841-516-000-0020-0360 "Duncan Galbraith"]{
            Duncan Antenupital Son to Neil Galbraith Maison in South
            Cuilgailtro and Elisabeth Stout their born
            26th June 1841 Registered 14th September 1846
        }
    }
    duncan1841-death {
        $sp-ref-link[d-1885-564-03-0396 0132 "Duncan Galbraith"]  Has a correction detailing the drowning accident.
    }
    rankin-marriage {
        $sp-ref-link[m-1868-564-03-0030 0015 "Duncan Galbraith" "Mary Rankin"]
    }
    census1871 {
        https://www.findmypast.com/transcript?id=GBC/1871/0022957220&expand=true
    }
    census1881 {
        https://www.findmypast.com/transcript?id=GBC/1881/0029499785&expand=true
    }

    marion1871-death {
        $sp-ref-link[d-1950-564-00-1017 0339 "Marion Shaw Galrbaith"]
    john1877-birth {
        $sp-ref-link[b-1877-564-02-0065 0022 "John Rankin Galbraith"]
    }
    john1877-death {
        $sp-ref[d-1881-564-01-0360 "John Rankin Galbraith"]
    }
}
