name Donald Galbreath
birth -date 'say 1795' -location colonsay
death -date 3-jan-1867 -location 'Elderslie Township, Bruce, Ontario, Canada'
tags Argyll:Colonsay Immigrant:Canada:Ontario
external {
    familysearch KH1J-4Z6
    findagrave 170628191
}
note {
    This may be two different Donalds.  While there are no conflicts, there is also no
    evidence that the Donald who married Rose McLean is the same one who married
    Annie McLugash.
}

body {
1841 Census, Argyllshire, Jura, Auchadaruch:$ref[census1841]
$csvtable{
First,Last,Age,Year
Donald,Galbreath,50,1791
Ann,Galbreath,40,1801
Angus,Galbreath,13,1828
Lauchlan,Galbreath,8,1833
Mary,Galbreath,5,1836
Sally,Galbreath,2,1839
Katrine,McLucas,47,1794
}

}

partner {
    name Rose McLean
    marriage -date 10-feb-1815 -location colonsay -ref mclean-marriage
    death -date 1821
    child {
        name John Galbreath
        baptism -date 26-feb-1816 -location colonsay -ref john1816-baptism
    }
    child galbreath-malcolm-1817-buie

    child {
        name Margat Galbreath
        baptism -date 12-sep-1819 -location colonsay -ref margat1819-baptism
    }
    child {
        name Cathrine Galbreath
        baptism -date 24-sep-1821 -location colonsay -ref cathrine1821-baptism
    }
}
partner {
    name Annie McLugash
    marriage -date 5-mar-1827 -location colonsay -ref mclugash-marriage
    birth -date 12-apr-1801 -location colonsay -ref mclugash-baptism
    death -date 3-jan-1867 -location 'Bruce County, Ontario, Canada'
    child galbreath-angus-1828-darrach
    child {
        name Lachlain Galbreath
        baptism -date 19-may-1833 -location colonsay -ref lachlain1833-baptism
        body { Present in 1861 census }
    }
    child {
        name Mary Galbreath
        baptism -date 20-jun-1836 -location colonsay -ref mary1836-baptism
        death -date 1823 -location 'Saskatchewan, Canada'
        partner {
            name Archibald McLugash
            marriage -date 27-jun-1855 -location colonsay -ref mary1836-marriage
        }
    }
    child {
        name Margaret Galbreath
        birth -date 'about 1840'
    }
    child {
        name Sarah Marion Galbreath
        birth -date 'about 1840'
        death -date 18-apr-1871 -location 'Arran, Bruce, Ontario, Canada'
    }
    child {
        name Anne Galbreath
        baptism -date 5-may-1844 -location colonsay -ref anne1844-baptism
        partner {
            name Donald McLean
            marriage -date 18-jun-1860 -location colonsay -ref anne1844-marriage
        }
    }
}
footnotes {
    mclean-marriage {
        $opr-ref[m-1815-539-020-0010-0018 "Donald Galbreath" "Rose McLean"]
    }
    john1816-baptism {
        $opr-ref[b-1816-539-020-0010-0011 "John Galbreath"]
    }
    margat1819-baptism {
        $opr-ref[b-1819-539-020-0010-0014 "Margat Galbreath"]
    }
    cathrine1821-baptism {
        $opr-ref[b-1821-539-020-0020-0005 "Cathrine Galbreath"]
    }
    mclugash-baptism {
        $opr-ref[b-1801-539-020-0010-0005 "Annie McLugas"]
    }
    mclugash-marriage {
        $opr-ref-link[m-1827-539-020-0020-0027 "Donald Galbraith" "Annie McLugash"]{
            March 5 | Donald Galbraith + Annie McLugash
        }
    }
    lachlain1833-baptism {
        $opr-ref[b-1833-539-020-0020-0009 "Lachlain Galbreath"]
    }
    mary1836-baptism {
        $opr-ref-link[b-1836-539-020-0020-0011 "Mary Galbreath"]{
            June 20 | Mary Daughter of Don'd and $i{blank}
        }
    }
    mary1836-marriage {
        $sp-ref-link[m-1855-539-02-0003 0002 "Mary Galbraith" "Archibald McLugash"]
    }
    anne1844-baptism {
        $opr-ref[b-1844-539-020-0020-0014 "Annie Galbreath"]
    }
    anne1844-marriage {
        $sp-ref-link[m-1860-539-02-0002 0001 "Ann Galbreath" "Donald McLean"]
    }

    census1841 {
        https://www.findmypast.com/transcript?id=GBC/1841/0016662488&expand=true
    }
}
