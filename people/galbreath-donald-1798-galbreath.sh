name Donald Galbreath
gender male
tags argyll:gigha

external {
    familysearch M1B4-5WZ
}
baptism -date 4-jan-1798 -location gigha -ref donald1798-baptism
death -date 28-aug-1860 -location gigha -ref donald1798-death

body {

From the 1827 Census of Gigha, in Achard:$ref[census1827]
$csvtable{
First,Last,Age,Birth
Donald, Galbraith, 29, 1798
Catherine, Galbraith, 27, 1800
}

}
note {
    Twins born, and both died during infancy.
}
partner {
    name Catherine Galbreath
    gender female
    parent galbreath-murdoch-1770-douglas
    baptism -date feb-1799 -location 'kilcalmonell and kilberry' 
    death -date 23-nov-1878 -location tarbert,argyll -ref catharine1799-death
    marriage -date 15-jul-1826 -location 'kilcalmonell and kilberry' -ref catharine1799-marriage

    child {
        name Lachlain Galbraith
        baptism -date 17-nov-1827 -location gigha -ref lachlain1827-baptism
    }
    child {
        name Katharine Galbraith
        baptism -date 13-apr-1829 -location gigha -ref katharine1829-baptism
        death -date 18-apr-1829 -location gigha -ref katharine1829-death
    }
    child {
        name Margarate Galbraith
        baptism -date 13-apr-1829 -location gigha -ref margaret1829-baptism
        death -date 25-apr-1829 -location gigha -ref margaret1829-death
    }
    child {
        name Margaret Galbraith
        baptism -date 25-jul-1830 -location gigha -ref margaret1830-baptism
        death -date 13-jun-1916 -location tarbert -ref margaret1830-death
        partner {
            name Donald McDonald
            marriage 13-dec-1849 -location gigha
        }
    }
    child {
        name Mary Galbreath
        baptism -date 03-apr-1831 -location gigha -ref mary1831-baptism
    }
    child galbraith-john-1833-wotherspoon

    child {
        name Catharine Galbraith
        baptism -date 5-jul-1835 -location gigha -ref catharine1835-baptism
    }   
}

footnotes {
    donald1798-baptism {
        $opr-ref[b-1798-537-000-0010-0013 "Donald Galbreath"]
    }
    donald1798-death {
        $sp-ref-link[d-1860-537-00-0003 0001 "Donald Galbraith"]
    }
    catharine1799-marriage {
        Recorded twice:
        $opr-ref[m-1826-516-000-0020-0220 "Donald Galbraith" "Cathrine Galbraith"];
        $opr-ref[m-1827-537-000-0020-0015 "Donald Galbreath" "Catharine Galbreath"]
    }
    catharine1799-death {
        $sp-ref-link[d-1878-535-00-0027 0009 "Catherine Galbraith"]
    }
    census1827 {
        From the 1827 Census of Gigha,
        https://www.gigha.org.uk/viewItem.php?id=8932
    }
    lachlain1827-baptism {
        $opr-ref[b-1827-537-000-0020-0016 "Lachlain Galbreath"]
    }

    katharine1829-baptism {
        $opr-ref[b-1829-537-000-0020-0019 "Katharine Galbreath"]
    }
    katharine1829-death {
        $opr-ref-link[d-1829-537-000-0020-0020 "Katharine Galbreath"]
    }
    margaret1829-baptism {
        $opr-ref[b-1829-537-000-0020-0019 "Margarat Galbreath"]
    }
    margaret1829-death {
        $opr-ref-link[d-1829-537-000-0020-0020 "Margaret Galbreath"]
    }
    margaret1830-baptism {
        $opr-ref[b-1830-537-000-0020-0021 "Margarat Galbreath"]
    }
    margaret1830-death {
        $sp-ref[d-1916-535-00-0014 "Margaret Galbraith"]
    }
    mary1831-baptism {
        $opr-ref[b-1831-537-000-0020-0025 "Mary Galbreath"]
    }
    catharine1835-baptism {
        $opr-ref[b-1835-537-000-0020-0030 "Catharine Galbreath"]
    } 
}
