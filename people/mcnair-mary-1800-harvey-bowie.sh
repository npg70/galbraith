name Mary McNair
gender female
tags Argyll:Campbeltown
external {
    familysearch 9M96-48T
    findagrave 166801432
}
birth -date 15-sep-1800
baptism -date 16-sep-1800 -location campbeltown -ref mary1800-baptism
death -date 21-nov-1863 -location campbeltown -ref mcnair-death
todo major {
    Lots of Marys to fix up.
}
body {
She is listed a a grocer in the census of 1841, living with her three sons.

Census of 1841, Longrow, Campbeltown, Argyllshire, Scotland:$ref[census-1841]
$csvtable{
First,Last,Age,Year
Mary,Harvey,35,1806
Robert,Harvey,11,1830
Archibald,Harvey,9,1832
John,Harvey,6,1835
}

Census of 1861, High Street, Campbeltown, Argyllshire, Scotland:$ref[census-1861]
$csvtable{
First,Last,Age,Year,Occupation
Archibald,Bowie,65,1796,Labouer
Mary,Bowie,60,1801,Wife of Labourer
Robert,Harvey,30,1831,Baker
Archibald,Harvey,29,1832,Fisherman
John,Harvey,26,1835,Baker
}
}
confused-with mcnair-mary-1780-johnston

partner harvey-robert-1791-mcnair
partner {
    name Archibald Bowie
    gender male
    birth -date 1797
    marriage -date 13-jun-1843 -location campbeltown -ref bowie-marriage
    death -date 03-apr-1867 -location campbeltown -ref bowie-death
    body {
        He was a gardener.
    }
}
footnotes {
    mary1800-baptism {
        $opr-ref-link[b-1800-507-000-0040-0206 "Mary McNair"]{
            Mary | Archibald McNair & Jean Mitchell Shiskan had a
            Dau Born 15th Baptized 16th named Mary
        }
    }
    mcnair-death {
        $sp-ref-link[d-1863-507-00-0149 0050 "Mary McNair"]
    }
    bowie-marriage {
        $opr-ref-link[m-1843-507-000-0060-0389 "Mary McNair" "Archibald Bowie"]{
        Archibald Bowie, Farmer, and Mary McNair
        relict of the deceased Robert Harvey,
        merchant, both of this parish, were married
        thirdteenth June 1843
        }
    }
    census-1841 {
        $elink[https://www.findmypast.com/transcript?id=GBC%2F1841%2F0016605612&tab=this]{FindMyPast}
    }
    census-1861 {
        $elink[https://www.findmypast.com/transcript?id=GBC%2F1861%2F0022161947&tab=this]{FindMyPast}
    }
    bowie-death {
        $sp-ref-link[d-1867-507-00-0045 0015 "Archibald Bowie"]{
            Widow of Mary McNair, Age 70.  Father: Donald Bowie, mother unlisted.
        }
    }
}
