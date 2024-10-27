name John Paterson
birth -date 'about 1764'
external {
    familysearch M1Y7-3ST
    findagrave 220562141
}
body {
    He was a cooper.
}
body {
Census of 1792, 86 Shore Street:
$csvtable{
Age,Name
?,John Paterson
24,Mary Fleeming
8,James Paterson
2, Janet Paterson
18,James Paterson
?, Mary Muir
}
}

partner {
    name Mary Fleeming
    birth -date 'about 1768' -ref fleeming-birth
    marriage -date 'about 1787' -ref fleeming-marriage
    child {
        name John Paterson
        birth -date 13-nov-1787
        baptism -date 14-nov-1787 -location campbeltown
    }
    child {
        name Janet Paterson
        birth -date 16-feb-1790
        baptism -date 17-feb-1790 -location campbeltown -ref janet1790-baptism
    }
    child {
        name Agnes Paterson
        birth -date 04-feb-1792
        baptism -date 06-feb-1792 -location campbeltown
    }
    child {
        name Mary Paterson
        baptism -date 28-jan-1794 -location campbeltown
    }
    child {
        name John Paterson
        baptism -date 30-mar-1796 -location campbeltown
    }
    child {
        name David Paterson
        birth -date 10-feb-1799
        baptism -date 13-feb-1799 -location campbeltown
    }
    child {
        name Janet Paterson
        birth -date 18-sep-1801
        baptism -date 20-sep-1801 -location campbeltown
        partner {
            name James Anderson
            marriage -date 26-jun-1818
        }
    }
    child {
        name Jean Paterson
        birth -date 15-mar-1805
        baptism -date 16-mar-1805 -location campbeltown
    }
    child {
        name James Paterson
        birth -date 08-aug-1807 
        baptism -date 10-aug-1807 -location campbeltown
        body {
            He married $child-link[galbraith-isabella-1808-paterson]{Isabella Galbraith}.
        }
    }
}
footnotes {
    fleeming-marriage {
        No record found.  1787 is estimated based on birth of first child.
    }
    fleeming-birth {
        The Census of 1792 lists her age as 24 (~1768).
    }
    janet1790-baptism {
        $opr-ref-link[b-1790-507-000-0020-0117 "Janet Paterson"]
        $blockquote{
            Janet | Lawful dau to John Paterson Cooper in Town + 
            Mary Fleeming was born 16th + Bapt 17th Feb [1790]
        }
    }
}
