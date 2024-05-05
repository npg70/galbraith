name John Galbreath
tags argyll:gigha 'daughtered out'
birth -date 'about 1760'
external {
    familysearch KZZS-S74
}
body {
    He was a schoolmaster.
}
body {
1827 Census of Gigha at Ardmeanish:$ref[census1827]
$csvtable{
First,Last,Age,Year
John,Galbraith,67,1769
Mary,Gilchrist,59,1768
Lachlan,Galbraith,30,1797
Mary,Galbraith,24,1803
}
}

partner {
    name Mary Gilchrist
    birth -date 'about 1768'
    marriage -date 15-feb-1785 -location 'killean and kilchenzie' -ref gilchrist-marriage

    child {
        name Anny Galbreath
        birth -date 'about 1792'
        body {
            ???
        }
    }
    child {
        name Lachlan Galbreath
        baptism -date 15-jun-1797 -location gigha -ref lachlan1797-baptism
        death -date 15-may-1867 -location 'killen and kilchenzie' -ref lachlan1797-death -note unmarried
    }
    child {
        name _____ Galbbreath
        baptism -date 08-dec-1799 -location gigha -ref unnamed1799-baptism
        death -date 16-dec-1867 -location gigha -ref unnamed1799-death
    }
    child {
        name Charles Galbreath
        baptism -date 03-may-1802 -location gigha -ref charles1802-baptism
        death -date 02-jun-1802 -location gigha -ref charles1802-death
    }
    child {
        name Mary Galbreath
        baptism -date 04-sep-1803 -location gigha -ref mary1803-baptism
    }
}
footnotes {
    census1827 {
        https://www.gigha.org.uk/viewItem.php?id=8932
    }
    gilchrist-marriage {
        $opr-ref-link[m-1785-519-000-0010-0311 "John Galbreath"]
    }
    lachlan1797-baptism {
        $opr-ref[b-1797-537-000-0010-0011 "Lachlan Galbreath"]
    }
    lachlan1797-death {
        $sp-ref-link[d-1867-519-00-0016 0006 "Lachlan Galbraith"]
    }
    unnamed1799-baptism {
        Indexed three times:
        $opr-ref[b-1799-537-000-0010-0011 "_____ Galbreath"];
        $opr-ref[b-1799-537-000-0010-0017 "_____ Galbreath"];
        $opr-ref[b-1799-537-000-0010-0018 "_____ Galbreath"]
    }
    charles1802-baptism {
        Indexed twice but with same identifier.
        $opr-ref[b-1802-537-000-0010-0023 "Charles Galbreath"]
    }
    charles1802-death {
        $opr-ref[d-1802-537-000-0010-0024 "Charles Galbreath"]
    }
    mary1803-baptism {
        $opr-ref[b-1803-537-000-0010-0026 "Mary Galbreath"]
    }
}
