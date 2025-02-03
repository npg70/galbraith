name Donald Galbraith
gender male
birth -date 'about 1815'

external {
    familysearch 9MMT-RLY
}
tags Argyll:Southend Immigrant:USA:Ohio:Washington

note {
    Family immigrated to Ohio between 1851 and 1861
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
}

body {
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
    gender female
    parent {
        name James Montgomery
        gender male
    }
    marriage -date 3-jun-1836 -location southend -ref marriage
    child {
        name Flora Galbraith
        baptism -date 11-feb-1837 -location southend -ref flora1837-baptism
        death -note dy -ref flora1837-death
    }
    child {
        name Archibald Galbraith
        baptism -date 6-mar-1838 -location southend -ref archibald1838-baptism
    }
    child {
        name James Galbraith
        birth -date 20-may-1840
        baptism -date 28-may-1840 -location southend -ref james1840-baptism
        death -note dy -ref james1840-death
    }
    child {
        name John Galbraith
        birth -date 20-feb-1842
        baptism -date 22-feb-1842 -location southend -ref john1842-baptism
    }

    child galbraith-alexander-1844-mcpherson

    child galbraith-margaret-1846-mcpherson

    child {
        name Donald Galbraith
        birth -date 9-jun-1849
        baptism -date 1-jul-1849 -location southend -ref donald1849-baptism
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
    flora1837-death {
        Death assumed based on absence in 1841 and 1851 Census.
    }
    archibald1838-baptism {
        $opr-ref[b-1838-532-000-0020-0093 "Archibald Galbraith"]
    }
    james1840-baptism {
        $opr-ref-link[b-1840-532-000-0020-0101 "James Galbraith"]{
            May 21 | James | Lawful son of Donald Galbraith Cottage
            Kildavie and Janet Montgomery was born 20th
            and baptized 28th current
        }
    }
    james1840-death {
        The death of James is inferred by his absense in the 1851 Census
    }
    john1842-baptism {
        $opr-ref-link[b-1842-532-000-0020-0107 "John Galbraith"]{
            Feb 21 | John | Lawful son of Donald Galbraith Cottages
            Kilevie[??] and Janet Montgomery was born
            20th and baptized 22 current [Feb]
        }
    }
    donald1849-baptism {
        $opr-ref-link[b-1849-532-000-0020-0123 "Donald Galbraith"]{
            June 30 | Donald lawful son of Donald Galbraith Labr.
            [???] and Janet Montgomery was born
            9th June and baptized 1st July
        }
    }
    census1841 {
        https://www.findmypast.com/transcript?id=GBC/1841/0016657570
    }
    census1851 {
        https://www.findmypast.com/transcript?id=GBC/1851/0019303086&expand=true
    }
}


