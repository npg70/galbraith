name Duncan Stewart
baptism -date 07-jul-1781 -location campbeltown -ref duncan1781-baptism
death -date jun-1848 -location peninver -ref duncan1781-death
burial -name 'Kilkerran Cemetery'
external {
    findagrave 158182855
    familysearch K8KH-96N
}
todo {
    8 children
}
parent {
    name William Stewart
    partner {
        name Mary McConnochy
    }
}

body {
$p{
He was a farmer in Caliburn, Campbeltown and then later to Peninver.
}

Census of 1841, Peninver, Campbeltown, Argyllshire, Scotland:$ref[census-1841]
$csvtable{
First,Last,Age,Year
Duncan, Stewart, 59, 1782
Margaret, Stewart, 56, 1785
William, Stewart, 25, 1816
Archibald, Stewart, 20, 1821
Dugald, Stewart, 15, 1826
Jean, Stewart, 12, 1829
}

}

partner {
    name Margaret Greenlees
    marriage -date 28-feb-1811 -location campbeltown -ref greenlees-marriage
    birth -date 1785
    death -date 1866 -location campbeltown

    child {
        name Mary Stewart
        gender female   
        birth -date 29-may-1816 -location campbeltown -ref mary-birth
        partner mcnair-john-1801-stewart
    }
}

footnotes {
    census-1841 {
        https://www.findmypast.com/transcript?id=GBC%2F1841%2F0016594975&expand=true&tab=this
    }
    duncan1781-baptism {
        $opr-ref-link[b-1781-507-000-0020-0054 "Duncan Stewart"]{
            Duncan | Lawful son to William Stewart & Mary McConnochy
            born 4th bapt 17th July 1781

       }
    }
    greenlees-marriage {
        $opr-ref-link[m-1811-507-000-0050-0092 "Duncan Stewart" "Margaret Greenlees"]{
            Duncan Stewart farmer Calliburn in this Parish and Margaret
            Greenlees of Killean Parish daughter to John Greenlees
            farmer [illegible] were married 28th February 1811.
        }
        $opr-ref-link[m-1811-519-000-0010-0365 "Duncan Stewart" "Peggy Greenlies"]{
            Feb | 28 | Duncan Stewart Calliburn parish of Campbeltown
            and Peggy Greenlies [illegible]
        }
    }
    duncan1781-death {
        Find a Grave, database and images (https://www.findagrave.com/memorial/158182855/duncan-stewart: accessed January 21, 2025), memorial page for Duncan Stewart (1780–Jun 1848), Find a Grave Memorial ID 158182855, citing Campbeltown Kilkerran Cemetery, Campbeltown, Argyll and Bute, Scotland; Maintained by LSP (contributor 46860931).
    }
    mary-birth {
        $opr-ref[b-1816-507-000-0040-0445 "Mary Stewart"]
    }
}

