name Duncan Galbraith
tags Argyll:Campbeltown
baptism -date 27-apr-1841 -location campbeltown -ref duncan1841-baptism
death -date 14-feb-1890 -location campbeltown -ref duncan1841-death
body {
    He was a tinsmith.
}
body {
Census 1881 @ New Quay Head Pensioners Row, Campbeltown, Argyllshire, Scotland:$ref[census1881]
$csvtable{
First,Last,Role,Age,Year,Birth,Occupation
Duncan,Galbraith,Head,39,1842,Campbeltown,Master Tin Smith
Jessie,Galbraith,Wife,34,1846,Paisley,-
Donald,Galbraith,Son,0,1881,Campbeltown,-
}
}

partner {
    name Jessie Mitchell
    marriage -date 24-jan-1878 -location campbeltown -ref mitchell-marriage

    child {
        name Donald Mitchell Galbraith
        birth -date 1880 -location campbeltown -ref donald1880-birth
        death -date 1880 -location campbeltown -ref donald1880-death
    }
    child {
        name George Mitchell Galbraith
        birth -date 13-nov-1881 -location campbeltown -ref george1881-birth
    }
    child {
        name Archibald Smith Galbraith
        birth -date 13-nov-1881 -location campbeltown -ref archibald1881-birth
    }
    child galbraith-malcolm-1884-simmons

}

footnotes {

    duncan1841-baptism {
        $opr-ref[b-1841-507-000-0070-0195 "Duncan Galbreath"]
    }
    duncan1841-death {
        $sp-ref-link[d-1890-507-00-0026 0009 "Duncan Galbraith"]
    }
    mitchell-marriage {
        $sp-ref-link[m-1878-507-00-0006 0003 "Duncan Galbraith" "Jessie Mitchell"]    
    }
    donald1880-birth {
        $sp-ref[b-1880-507-00-0334 "Donald Mitchell Galbraith"]
    }
    donald1880-death {
        $sp-ref[d-1880-507-00-0082 "Donald Mitchell Galbraith"]
    }
    george1881-birth {
        $sp-ref-link[b-1881-507-00-0348 0116 "George Mitchell Galbraith"]
    }
    archibald1881-birth {
        $sp-ref-link[b-1881-507-00-0349 0117 "Archibald Smith Galbraith"]
    }
    census1881 {
        https://www.findmypast.com/transcript?id=GBC/1881/0029347276&expand=true
    }    
}
