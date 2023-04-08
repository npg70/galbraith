name John Galbreath
birth -date 1788 
death -date 18-sep-1870 -location Skipness,Argyll -ref john1788-death 
external {
    familysearch 9M9X-D8G
}
tags 'Kilcalmonell and Kilberry'

body {
From death record, the parents where Archibald Galbreath and Mary (Marg?) McIntyre.

Maybe -- Census of 1861, Scotland, Argyllshire, Saddell, Saddell Village:$ref[census1861]
   $csvtable{
First,Last,Role,Age,Date,Place,Job
John,Galbreath,Head,68,1793,Kilbery,Ag Lab
Mary,Galbreath,Wife,70,1791,Agyllshire,
}

}

partner {
    name Mary Sinclair
    marriage -date 28-apr-1816 -location 'Kilcalmonell and Kilberry' -ref sinclair-marriage

    child {
        name Neil Galbraith
        birth -date 'about 1817'
        death -date 23-mar-1874 -location Saddell -ref neil1817-death
        partner {
            name Isabella McMillan
        }
    }
    child galbraith-donald-1819-smith
    child {
        name Peggy Galbreath
        birth -date 7-nov-1821
        baptism -location 'Kilcalmonell and Kilberry' -ref peggy1821-baptism
    }
    child {
        name Peter Galbraith
        birth -date 10-jul-1824
        baptism -date 24-jul-1824 -location 'Kilcalmonell and Kilberry' -ref peter1824-baptism
    }
    child {
        name Mary Galbraith
        birth -date 7-jan-1827
        baptism -location 'Kilcalmonell and Kilberry' -ref mary1827-baptism
    }
    child {
        name Marion Galbraith
        birth -date 14-jul-1829
        baptism -location 'Kilcalmonell and Kilberry' -ref marion1829-baptism
    }
    child {
        name John Galbraith
        birth -date 15-feb-1833 -location 'North Kilberry'
        baptism -date 12-mar-1833 -location 'Kilcalmonell and Kilberry' -ref john1833-baptism
    }
}

footnotes {
    john1788-death {
        $sp-ref-link[d-1870-531-02-0005 0002 "John Galbraith"]
    }
    sinclair-marriage {
        $opr-ref-link[m-1816-516-000-0010-0415 "John Galbreath" "Mary Sinclair"]{
           Clachan April 28 1816 - The which day were booked in order to
           marriage John Galbreath + Mary Sinclair both in this parish
           of Kilberry + $i{illegible}
        }
    }
    neil1817-birth {
        Birth appears to be unrecorded, however confirmed in Census of 1851, and
        death record, below.
    }
    neil1817-death {
        $sp-ref-link[d-1874-531-01-0001 0001 "Neil Galbraith"]
    }
    peggy1821-baptism {
        $opr-ref[b-1821-516-000-0020-0021 "Peggy Galbreath"]
    }
    peter1824-baptism {
        $opr-ref-link[b-1824-516-000-0020-0297 "Peter Galbraith"]{
            Peter lawful son to John
            Galbraith + Mary Sinclair in
            Gontannariah[?] born the 10th 
            July 1824  Registed the 24th July 1824
        }
    }
    mary1827-baptism {
        $opr-ref[b-1827-516-000-0020-0305 "Mary Galbraith"] 
    }
    marion1829-baptism {
        $opr-ref[b-1829-516-000-0020-0315 "Marion Galbraith"]
    }
    john1833-baptism {
        $opr-ref-link[b-1833-516-000-0020-0333 "John Galbraith"]{
            John lawful son of John Galbraith Fisherman
            North Kilberry and Mary Sinclair there born
            15th Febry 1833 RegD 12th March 1833
        }
    }
    census1861 {
        https://www.findmypast.com/transcript?id=GBC%2F1861%2F0022530617
    }
}
