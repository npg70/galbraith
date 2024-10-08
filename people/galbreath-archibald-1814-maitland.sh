name Archibald Galbreath
birth -date 'about 1814' -location Gigha
external {
    familysearch KH8D-G3C
}
tags Argyll:Campbeltown 'Dead End'

body {
1841 Census, Store House New Quay Head, Campbeltown$ref[census-1841]
$csvtable{
Name,Age,Year
Archibald Galbreath,25,1816
John Galbreath,2,1839
John Maitland,55,1786
Catharine McLiver,45,1786
}


1851 Census, Shore Street East Side, Campbeltown,  507-11-16$ref[census-1851]
$csvtable{
Name,Relation,Age,Year,Birthplace,Occupation
Archibald Galbreath,Head,38,1813,Gigha,Fisherman
Elizabeth Galbreath,Wife,36,1815,Campbeltown,-
John Galbreath,Son,12,1839,Campbeltown,Scholar
Flora Galbreath,Dau,7,1844,Campbeltown,Scholar
}
}

note {
    No records found after 1851 census.
}
note {
    Maybe parents in Gigha
}

partner {
    name Elizabeth Maitland
    birth -date 'about 1815' -location campbeltown
    marriage -date 17-nov-1835 -location campbeltown -ref maitland-marriage
    child {
        name John Galbreath
        birth -date 21-jun-1839
        baptism -date 26-mar-1840 -location campbeltown -ref john1840-baptism
    }
    child {
        name Flora Galbreath
        baptism -date 6-nov-1843 -location campbeltown -ref flora1843-baptism
    }
}

footnotes {
    maitland-marriage {
        $opr-ref-link[m-1835-507-000-0060-0325 "Archibald Galbreath" "Elizabeth Maitland"]{
            Archibald Galbreath laborer and Elizabeth Maitland
            both of this Parish were married 16th Nov 1835
        }
    }
    john1840-baptism {
        $opr-ref-link[b-1840-507-000-0070-0169 "John Galbreath"]{
            John - Lawful son of Archibald Galbreath Fisherman
            and Elizabeth Maitland was born the
            twenty-first June 1839 and Bapt 26th March
        }
    }
    flora1843-baptism {
        $opr-ref-link[b-1843-507-000-0070-0249 "Flora Galbreath"]{
            Flora - Lawful daughter of of Archd Galbreath Seaman
            and Elizabeth Maitland was born fiftheenth October
            and Bapt 6th Nov 1843.
        }
    }
    census-1841 {
        https://www.findmypast.com/transcript?id=GBC%2F1841%2F0016593065&expand=true&tab=this
    }
    census-1851 {
        https://www.findmypast.com/transcript?id=GBC%2F1851%2F0019261122&expand=true&tab=this
    }
}
