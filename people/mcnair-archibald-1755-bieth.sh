name Archibald McNair
gender male
tags 'Argyll:killean and kilchenzie'
birth -date 1755 -location campbeltown
todo {
    familysearch is a mess.
}
note {
    Given marriage and baptisms of children, they were likely both from Killean. Birth records
    are unlikely to be found since record keeping started around 1762 in Killean.
}
note {
    Confused with $child-link[mcnair-archibald-1753-mitchell]{Archibald McNair}, b. 1753, m. $i{Jean} Mitchell
}
note {
    Confused with $child-link[mcnair-archibald-1759-mitchell]{Archibald McNair}, b. 1759, m. $i{Margaret} Mitchell
}

body {
    He was a sailor, according to his daughter Mary's death certificate.
}
body {
Census of 1792, Parish of Killean, Belochgoichan:
$csvtable{
Age,Year,Men,Women,Children,Role
37,1755,Archibald McNair, , ,Head
36,1756,                , Mary Beith, , Spouse
70,1722,                , Cathrin McCorvie, , Mother or MIL?
 6,1786,                ,                   , Angus McNair, Daughter
 2,1790,                ,                   , Mary McNair, Daughter
}

Angus is "Anny"  (or vice versa).
}
partner {
   name Mary Bieth
   gender female
   birth -date 1756
   marriage -date 15-mar-1785 -location 'killean and kilchenzie' -ref bieth-marriage

   child {
        name Anny McNair
        gender female
        baptism -date 01-may-1786 -location 'killean and kilchenzie'  -ref anny1786-baptism
   }
   child {
        name Cathrine McNair
        gender female
        baptism -date 11-jul-1788 -location 'killean and kilchenzie'  -ref cathrine1788-baptism
   }
   child {
        name Mary McNair
        gender female
        baptism -date 18-may-1790 -location 'killean and kilchenzie'  -ref mary1790-baptism
        death -date 14-mar-1867 -location 'killean and kilchenzie' 
        partner {
            name Archibald McNiel
            gender male
        }
   }
   child {
        name Margaret McNair
        gender female
        baptism -date 19-dec-1792 -location 'killean and kilchenzie'  -ref margaret1792-baptism
   }
   child {
        name Edward Iver McNair
        gender male
        baptism -date 10-nov-1796 -location 'killean and kilchenzie'  -ref iver1796-baptism
        death -date 23-feb-1861 -location campbeltown -ref iver1796-death
        partner galbraith-margaret-1795-mcnair
   }
}

footnotes {
    bieth-marriage {
        $opr-ref-link[m-1785-519-000-0010-0311 "Archibald McNair" "Mary Bieth"]{
            Archd McNair and Mary Bieth both of this Parish 
            Contracted married 15th March
       }
    }
    anny1786-baptism {
        $opr-ref[b-1786-519-000-0010-0042 "Anny McNair"]
    }
    cathrine1788-baptism {
        $opr-ref[b-1788-519-000-0010-0054 "Cathrine McNair"]
    }
    mary1790-baptism {
        $opr-ref[b-1790-519-000-0010-0066 "Mary McNair"]
    }
    margaret1792-baptism {
        $opr-ref[b-1792-519-000-0010-0088 "Margaret McNair"]
    }
    iver1796-baptism {
        $opr-ref[b-1796-519-000-0010-0119 "Iver McNair"]
    }
    iver1796-death {
        $sp-ref-link[d-1861-507-00-0025 0009 "Edward McNair"]
    }
}
