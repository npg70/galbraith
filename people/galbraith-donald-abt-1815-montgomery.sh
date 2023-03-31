name Donald Galbraith
birth -date 'about 1815'

external {
    familysearch 9MMT-RLY
}

note {
    No additional information after 1851 census.  Perhaps immigrated?
}

body {

1841 Scotland, Argyll, Southend, Kildanie:$ref[census1841]

$csvtable{
First,Last,Age,Birth,Place
Donald,Galbreath,35,1806,Argyllshire
Janet,Galbreath,25,1816,Argyllshire
Archibald,Galbreath,3,1838,Argyllshire
James,Galbreath,1,1840,Argyllshire
}

1851 Scotland, Argyll, Southend, Moneroy Village:$ref[census1851] 
$csvtable{
First,Last,Role,Married,Age,Year,Birth,Occupation
Janet,Galbraith,Head,Married,37,1814,Campbeltown,Agr lab wife
Archibald,Galbraith,Son,-,13,1838,Southend,
John,Galbraith,Son,-,8,1843,Southend,
Alexander,Galbraith,Son,-,6,1845,Southend,
Margaret, Galbraith,-,-,4,1847,Southend,
Donald,Galbraith,-,-,1,1850,Southend,
}

}

partner {
    name Janet Montgomery
    marriage -date 3-jun-1836 -location southend -ref marriage
    child {
        name Flora Galbraith
        baptism -date 11-feb-1837 -location southend -ref flora1837-baptism
    }
    child {
        name Archibald Galbraith
        baptism -date 6-mar-1838 -location southend -ref archibald1838-baptism
    }
    child {
        name James Galbraith
        baptism -date 20-may-1840 -location southend -ref james1840-baptism
    }
    child {
        name John Galbraith
        baptism -date 20-feb-1842 -location southend -ref john1842-baptism
    }
    child {
        name Margaret Galbraith
        baptism -date 15-feb-1846 -location southend -ref margaret1846-baptism
    }
    child {
        name Donald Galbraith
        baptism -date 9-jun-1849 -location southend -ref donald1849-baptism
    }
}
footnotes {
    marriage {
        $opr-ref-link[m-1836-532-000-0020-0162 "Donald Galbraith" "Janet Montgomery"]{
            June 3 | Donald Galbraith Servant Bruenericken and
            Janet Montgomery daughter of James Montgomery 
            Cottages Ariseauach were married 9th $i{other illegible}        
        }
    }
    flora1837-baptism {
        $opr-ref[b-1837-532-000-0020-0089 "Flora Galbraith"]
    }
    archibald1838-baptism {
        $opr-ref[b-1838-532-000-0020-0093 "Archibald Galbraith"]
    }
    james1840-baptism {
        $opr-ref[b-1840-532-000-0020-0101 "James Galbraith"]
    }
    john1842-baptism {
        $opr-ref[b-1842-532-000-0020-0107 "John Galbraith"]
    }
    margaret1846-baptism {
        $opr-ref[b-1846-532-000-0020-0117 "Margaret Galbraith"]
    }
    donald1849-baptism {
        $opr-ref[b-1849-532-000-0020-0123 "Donald Galbraith"]
    }
    census1841 {
        https://www.findmypast.com/transcript?id=GBC/1841/0016657570
    }
    census1851 {
        https://www.findmypast.com/transcript?id=GBC/1851/0019303086&expand=true
    }
}


