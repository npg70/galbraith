name John Galbreath
tags Argyll:Gigha
birth -date 'say 1770'
external {
    familysearch M1B4-HNT
}
note {
    Spouse was baptized as Marrion/Merror but used the
    angelization of Sarah.
}
body {
Isle of Gigha Census 1827:$ref[census1827]
$csvtable{
First,Last,Age,Year
John, Galbraith, 60, 1767
Donald, Galbraith, 29, 1798
John, Galbraith, 27, 1800
Catharine, Galbraith, 22, 1805
Mary, Galbraith, 17, 1810
Sarah,-, 15, 1812
Malcolm,-, 12, 1815
James,-, 9, 1818
}
}

partner {
    name Marion Sarah Galbreath
    marriage -date 17-jun-1795 -location gigha -ref marion-marriage

    child {
        name Archibald Galbreath
        baptism -date 19-jul-1796 -location gigha -ref archibald1796-baptism
        death -date 01-apr-1881 -location dunoon -ref archibald1796-death
        partner {
            name Janet Morrison
            marriage -date 16-may-1826 -location greenock -ref morrison-marriage
        }
    }

    child {
        name Donald Galbreath
        baptism -date 31-mar-1798 -location gigha -ref donald1798-baptism
        death -date 10-jan-1879 -location gigha -ref donald1798-baptism
    }

    child galbreath-john-1799-mcquilkan

    child {
        name Margaret Galbreath
        baptism -date 28-apr-1803 -location gigha -ref margaret1803-baptism
        death -date 28-oct-1886 -location gigha -ref margaret1803-death
        partner {
            name John Smith
        }
    }

    child {
        name Catharine Galbreath
        baptism -date 06-may-1805 -location gigha -ref catharine1805-baptism
    }
    child {
        name Hector Galbreath
        baptism -date 18-aug-1807 -location gigha -ref hector1808-baptism
    }
    child galbraith-mary-1810-templeton

    child {
        name Merran Galbreath
        baptism -date 24-may-1812 -location gigha -ref merran1812-baptism
        death -date 20-sep-1864 -location gigha -ref merran1812-death
        partner {
            name John Willie?
            body {
                May be John McQuilkan
            }
        }
    }
    child galbreath-malcolm-1815-galbreath

    child galbreath-james-1817-wotherspoon
}
footnotes {
    marion-marriage {
        $opr-ref[m-1795-537-000-0010-0008 "John Galbreath" "Marion Galbreath"]
    }
    archibald1796-baptism {
        $opr-ref[b-1796-537-000-0010-0010 "Archibald Galbreath"]
    }
    archibald1796-death {
        $sp-ref-link[d-1881-510-00-0039 0013 "Archibald Galbraith"]
    }
    morrison-marriage {
        $opr-ref-link[m-1826-564-003-0080-0114 "Archibald Galbraith" "Janet Morison"]
    }
    donald1798-baptism {
        $opr-ref[b-1798-537-000-0010-0013 "Donald Galbreath"]
    }
    donald1798-baptism {
        $sp-ref-link[d-1879-537-00-0001 0001 "Donald Galbraith"]
    }
    margaret1803-baptism {
        $opr-ref[b-1803-537-000-0010-0025 "Margarate Galbreath"]
    }
    margaret1803-death {
        $sp-ref-link[d-1886-537-00-0011 0004 "Margaret Galbraith"]
    }
    catharine1805-baptism {
        $opr-ref-link[b-1805-537-000-0010-0029 "Catharine Galbreath"]
    }
    hector1808-baptism {
        Indexed twice, 
        $opr-ref[b-1807-537-000-0010-0034 "Hector Galbreath"];
        $opr-ref[b-1808-537-000-0010-0037 "Hector Galbreath"]
    }
    merran1812-baptism {
        $opr-ref-link[b-1812-537-000-0010-0048 "Marran Galbreath"]
    }
    merran1812-death {
        $sp-ref-link[d-1864-537-00-0006 0002 "Sarah Galbreath"]
    }
}

