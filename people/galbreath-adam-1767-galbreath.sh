name Adam Galbreath
tags argyll:gigha
external {
    familysearch KZLH-6VP
}
birth -date 'about 1767'
todo { Needs references }

body {
From the 1827 Census of Gigha at Drumyeonmore:$ref[census1827]
$csvtable{
First,Last,Age,Birth
Adam, Galbraith, 60, 1767
Flora, Galbraith, 52, 1775
Peter, Galbraith, 27, 1800
Catherine, Galbraith, 25, 1802
Neil, Galbraith, 21, 1806
Peggy, Galbraith, 19, 1808
Elizabeth, Galbraith, 16, 1811
Lachlan, Galbraith, 14, 1813
Nelly, Galbraith, 11, 1816
Malcolm, Galbraith, 7, 1820
}

}

partner {
    name Flora Galbreath
    birth -date 'about 1773'
    marriage -date 07-dec-1796 -location gigha -ref flora-marriage
    death -date 27-nov-1856 -location gigha -ref flora-death
    body {
        She was the daughter of $child-link[galbreath-neil-1750-galbreath]{Neil Galbreath} and Catherine Galbreath.
    }
    child galbreath-peter-1798-galbreath
    child {
        name Catharine Galbreath
        baptism -date 06-jun-1800 -location gigha -ref catharine1800-baptism
        death -date 06-aug-1880 -location gigha -ref catharine1800-death -note unmarried
    }
    child {
        name John Galbreath
        baptism -date 14-nov-1802 -location gigha -ref john1802-baptism
    }
    child {
        name Neil Galbreath
        baptism -date 09-may-1805 -location gigha -ref neil1805-baptism
    }
    child {
        name Margarate Galbreath
        baptism -date 01-aug-1807 -location gigha -ref marg1807-baptism
    }
    child {
        name Margarate Galbreath
        baptism -date 01-aug-1808 -location gigha -ref marg1808-baptism
    }
    child {
        name Isobel Galbreath
        baptism -date 12-nov-1809 -location gigha -ref isobel1809-baptism
        death -date 27-mar-1876 -location gigha -ref isobell1809-death
        body {
            Unmarried.
        }
    }
    child {
        name Lachlan Galbreath
        baptism -date 28-feb-1812 -location gigha -ref lachlan1812-baptism
        death -date 18-jun-1875 -location gigha -ref lachlan1812-death
        body {
            Unmarried.
        }
    }
    child {
        name Nelly Galbreath
        baptism -date 19-jan-1815 -location gigha -ref nelly1815-baptism
    }
    child {
        name Malcolm Galbreath
        baptism -date 10-jun-1819 -location gigha -ref malcolm1819-baptism
        death -date 07-apr-1900 -location gigha -ref malcolm1819-death
        body {
            Unmarried.
        }
    }
}
footnotes {
    flora-marriage {

    }
    flora-death {
        $sp-ref-link[d-1856-537-00-0005 0002 "Flora Galbreath"]
    }    
    catharine1800-baptism {
        $opr-ref[b-1800-537-000-0010-0018 "Catherine Galbreath"]
    }
    catharine1800-death {
        $sp-ref-link[d-1880-537-00-0008 0003 "Catherine Galbraith"]
    }
    john1802-baptism {
        $opr-ref[b-1802-537-000-0010-0023 "John Galbreath"]
    }
    neil1805-baptism {
        $opr-ref-link[b-1805-537-000-0010-0029 "Neil Galbreath"]
    }
    marg1807-baptism {
        $opr-ref[b-1807-537-000-0010-0034 "Margarate Galbreath"], father listed as "Adom"
    }
    marg1808-baptism {
        $opr-ref[b-1808-537-000-0010-0037 "Margarate Galbreath"], father listed as "Adom"
    }
    isobel1809-baptism {
        $opr-ref[b-1809-537-000-0010-0042 "Isobel Galbreath"]
    }
    isobell1809-death {
        $sp-ref-link[d-1876-537-00-0001 0001 "Isabella Galbraith"]
    }
    lachlan1812-baptism {
        $opr-ref-link[b-1812-537-000-0010-0048 "Lachlan Galbreath"]
    }
    lachlan1812-death {
        $sp-ref-link[d-1875-537-000-0009 0003 "Lachlan Galbraith"]
    }
    nelly1815-baptism {
        $opr-ref[b-1815-537-000-0010-0053 "Nelly Galbreath"]
    } 
    malcolm1819-baptism {
        $opr-ref[b-1819-537-000-0010-0062 "Malcom Galbreath"]
    }
    malcolm1819-death {
        $sp-ref-link[d-1900-537-00-0007 0003 "Malcolm Galbraith"]
    }
    census1827 {
        From the 1827 Census of Gigha,
        https://www.gigha.org.uk/viewItem.php?id=8932
    }
}
