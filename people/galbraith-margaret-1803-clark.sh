name Margaret Galbreath
tags Argyll:Campbeltown
baptism -date 19-jan-1803 -location campbeltown -ref margaret1803-baptism

external {
    familysearch KHY3-J5F
}
todo {
    Daughters not listed in Family Search in 2024.
}

note {
Margaret's last child was in 1842 and her husband is listed as a widow in 1851, implied she died between 1842 and 1851.
}

note {
William appears in the 1851 Census but not in statutory death records, impling he died between 1851 and 1855.
}

partner {
    name William Clark
    birth -date 'about 1789' -location campbeltown
    marriage -date 3-apr-1826 -location campbeltown -ref clark-marriage

    child {
        name Peter Clark
        birth -date 07-oct-1827
        baptism -date 10-oct-1827 -location campbeltown
    }
    child {
        name William Clark
        birth -date 12-apr-1829
        baptism -date 19-apr-1829 -location campbeltown
    }
    child {
        name Margaret Clark
        birth -date 21-aug-1831 -location campbeltown -ref marg1831
    }
    child {
        name Catharine Clark
        birth -date 22-apr-1842 -location campbeltown
    }
    child {
        name Margaret Clark
        birth -date 22-apr-1842 -location campbeltown
    }
}

body {
Census of 1841 at Longrow, Campbeltown, Argyllshire, Scotland:$ref[census-1841]
$csvtable{
First,Last,Age,Year,Occupation
William,Clark,55,1786,Malster
Margaret,Clark,36,1806,-
Peter,Clark,13,1828,-
William,Clark,11,1830,-
}
}
body {
Census of 1851 at Longrow East Side, Campbeltown, Cantyre, Argyllshire, Scotland:$ref[census-1851]
$csvtable{
First,Last,Status,Age,Year,Place,Occupation
William,Clark,Widow,62,1789,Campbeltown,Malster
Margaret,Clark,-,8,1843,Campbeltown,Scholar
}
}


footnotes {
    margaret1803-baptism {
        $opr-ref-link[b-1803-507-000-0040-0243 "Margaret Galbreath"]{
            Margaret | Lawful Daughter to Archibald Galbreath Merchant in Campbeltown
            and Jean Corbert. Born 16th baptized 19th January 1803. Named Margaret.
        }
    }

    marg1831 {
        1831 CLARK, MARGARET (Old Parish Registers Births 507/ Campbeltown) Page 226 of 478, indexed under William Clark and Mary Galbreath.  It appears ScotlandsPeople mistranscribed 'Marg' as "Mary".
        $blockquote{
            Margaret | Lawful Daughter of William Clark Malster +
            Margt Galbreath was born [??] + Bapt 21 Aug 1831. Margt.
       }
    }
    clark-marriage {
        $opr-ref-link[m-1826-507-000-0060-0273 "Margaret Galbreath" "William Clark"]{
            William Clark malster + Margt Galbreath
            both of this Parish were married 3 April 1826
        }
    }

    census-1841 {
        https://www.findmypast.com/transcript?id=GBC/1841/0016601132
    }
    census-1851 {
        https://www.findmypast.com/transcript?id=GBC/1851/0019257980        
    }
}
