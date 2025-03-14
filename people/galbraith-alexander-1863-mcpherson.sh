name Alexander Galbraith
gender male
tags Glasgow Argyll:Dunoon
external {
    familysearch 9Q3N-7FL
}
body {
Census of 1891: Bellmont Place, Dunoon & Kilmun, Dunoon, Argyllshire, Scotland:$ref[census1891]
$csvtable{
First,Last,Role,Status,Age,Year,Place,Job
Duncan,McPherson,Head,Widower,74,1817,Argyll,Retired Teacher
Alexander,Galbraith,Son in Law,Married,27,1864,Glasgow,Coachman
Elizabeth,Galbraith,Daughter,Married,27,1864,-,Greenock
James,McPherson,Son,Unmarried,24,1867,Greenock,Seaman
}
}

birth -date 13-aug-1863 -location glasgow -ref alexander1863-birth
partner {
    name Elizabeth Davie MacPherson
    gender female
    parent {
        name Duncan McPherson
        partner {
            name Elizabeth Wilson Davie
        }
    }
    birth -date 30-jul-1860 -location greenock
    death -date 1917 -location paisley -ref macpherson-death
    marriage -date 24-sep-1890  -location Dunoon -ref macpherson-marriage

    child {
        name Alexander Davie Galbraith
        birth -date 1891 -location dunoon -ref alexander1891-birth
        death -date 1952 -location hillhead -ref alexander1891-death
    }
    child {
        name Malcolm Hendry Galbraith
        birth -date 17-jul-1894 -location dunoon -ref malcolm1894-birth
        death -date 13-mar-1955 -location greenock -ref malcolm1894-death -note unmarried.
    }
}

footnotes {
    alexander1863-birth {
        $sp-ref-link[b-1863-644-09-1178 0393 "Alexander Galbraith"]
    }
    macpherson-marriage {
        $sp-ref-link[m-1890-510-01-0022 0011 "Alexander Galbraith" "Elizabeth Davie MacPherson"]
    }

    alexander1891-birth {
        $sp-ref[b-1891-510-01-0090 "Alexander Davie Galbraith"]
    }
    alexander1891-death {
        $sp-ref[d-1891-644-13-0908 "Alexander Davie Galbraith"]
    }
    malcolm1894-birth {
        $sp-ref-link[b-1894-510-01-0091 0031 "Malcolm Hendry Galbraith"]
    }
    malcolm1894-death {
        $sp-ref-link[d-1955-564-00-0249 0083 "Malcolm Henry Galbraith"]
    }

    census1891 {
        https://www.findmypast.com/transcript?id=GBC/1891/0035194337&expand=true
    }
}
