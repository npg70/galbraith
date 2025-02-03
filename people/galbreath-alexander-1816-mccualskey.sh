name Alexander Galbreath
gender male
baptism -date 20-apr-1816 -location southend -ref alexander1816-baptism
death -date 1898 -location Washington,Ohio,USA
burial -name 'Rockland Cemetery' -location 'Belpre,Washington,Ohio,USA'
tags Argyll:Southend Immigrant:USA:Ohio:washington
external {
    familysearch 9NN7-9JW
    findagrave 5024675
}
note {
    FindAGrave page says Alexander was born in Lochgilphead without source, and  which makes no sense.
}
todo {
    DNA match and direct male descendant on $elink[https://www.ancestry.com/profile/01543c3f-0001-0000-0000-000000000000?compareToTestId=3C22E008-8305-42DB-B065-DA41ACDCAA23]{Ancestry}
}

body {
    In the 1841 Census he is listed as a laborer on farm of the Brown family in  Machremore, Southend.$ref[census1841]
}
body {
Census of 1851, Saddel Street, Campbeltown, Cantyre, Argyllshire, Scotland:$ref[census1851]
$csvtable{
First,Last,Role,Status,Age,Year,Place,Job
Alexander,Galbraith,Head,Married,35,1816,Southend,Malt Man
Elizabeth,McCulcy,Wife,Married,36,1815,Campbeltown,-
Archibald,Galbraith,Son,-,5,1846,Campbeltown,-
Alexander,Galbraith,Son,-,0,1851,Campbeltown,-
Margaret,Reside,Mother in Law,Widow,64,1787,Campbeltown,x Weaver's Wife
Edward,McCulcy,Brother in Law,Married,34,1817,Campbeltown,Joiner
Agness,Hall,Sister in Law,Married,33,1818,Glasgow,-
Dugald,McCulcy,-,-,0,1851,Campbeltown,-
}
}

partner {
    name Elisabeth McCualskey
    gender female
    parent {
        name Alexander McCualskey
        birth -date 15-aug-1772 -location southend
        death -date 20-feb-1847 -location campbeltown
        partner {
            name Margaret Raeside
            birth -date 26-dec-1786 -location campbeltown
            death -date 20-apr-1872 -location campbeltown
        }
    }
    marriage -date 6-Feb-1845 -location campbeltown -ref marriage
    birth -date 6-may-1815
    baptism -date 8-may-1815 -location campbeltown  -ref eliz-baptism
    death -date 1904 -location "Parkersburg,Wood,West Virginia,USA"
    burial -name 'Rockland Cemetery' -location 'Belpre,Washington,Ohio,USA'

    child {
        name Archibald Galbreath
        birth -date  13-jan-1846
        baptism -date 20-jan-1846 -location southend -ref archibald1846-baptism
    }
    child {
        name Margaret Galbreath
        birth -date 20-sep-1848
        baptism -date 15-oct-1848 -location campbeltown -ref margaret1848-baptism
    }
    child {
        name Alexander Galbreath
        birth -date 'about 1850' -ref alex1850-birth
        death -date 1-July-1925 -location "Parkersburg,Wood,West Virginia,USA"
        partner {
            name Pheobe Jane Burnfield
        }
    }    
    child {
        name Helen Galbreath
        birth -date 'about 1852' -ref helen-birth
    }
    child {
        name Margaret Elizabeth Galbreath
        birth -date 1858 -location veto,ohio,usa
        death -date 12-jul-1937 -location "wood,west virginia,usa"
    }
}

footnotes {
    alexander1816-baptism  {
        $opr-ref-link[b-1816-532-000-0010-0102 "Alexander Galbreath"]{
            April 20 | Alexd lawful son of Archd Galbreath Mill of Machimore
        }
    }
    marriage {
        $opr-ref-link[m-1845-532-000-0020-0167 "Alexander Galbraith" "Elisabeth McCualskey"]{
            May 30 Alexander Galbraith Servant Marchrebeg and
            Elisabeth McCualskey of Campbeltown Parish were married 2 ? June.
        }
        $opr-ref-link[m-1845-507-000-0060-0405 "Alexander Galbreath" "Elizabeth McCualesky"]{
           Alexander Galbreath Workman Southend Parish and Elizabeth McCualsky
           daughter of Alexander McCualsky Weaver Campbeltown were married
           2nd June 1845
       }
    }
    archibald1846-baptism {
        $opr-ref-link[b-1846-532-000-0020-0117 "Archibald Galbraith"]{
           Archibald lawful son of Alexander Galbraith
           Laborer Machebeg[?] and Elizabeth McCualskey
           was born 13th and bapt 20th $i{Jan 1846}
        }
    }
    margaret1848-baptism {
        $opr-ref-link[b-1848-507-000-0070-0342 "Margaret Galbreath"]{
            Margaret | Lawful daughter of Alexander Galbreath Farm Servant
            Elizabeth Culasky born 20th September and Baptized
            15th October 1848
        }
    }
    census1841 {
        https://www.findmypast.com/transcript?id=GBC%2F1841%2F0016659062&tab=this
    }
    census1851 {
        https://www.findmypast.com/transcript?id=GBC%2F1851%2F0019255793&tab=this
    }
    eliz-baptism {
        $opr-ref-link[b-1815-507-000-0040-0429 "Elizbeth MacCullasky"]{
            Elizabeth | Lawful Daughter to Alexander MacCullasky Weaver 
            Campbeltown and Margaret Ryside born 6th bapt 8th May 1815
            Named Elizabeth
        }
    }
    alex1850-birth {
        Not in OPR
    }
    helen-birth {
        Not in OPR
    }
}
