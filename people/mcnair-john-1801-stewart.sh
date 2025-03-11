name John McNair
gender male
tags Argyll:Campbeltown endofline argyll:smerby

external {
    familysearch M1T6-PXV
}

birth -date 23-apr-1801
baptism -date 26-apr-1801 -location campbeltown
death -date 25-sep-1874 -location campbeltown

body {
    He was farm manager, and from Census records lived apart from his family.$ref[census-1851]$ref[census-1861]  Later he became a road contractor.$ref[census-1871]  None of children got married or had children.
}

body {
Census of 1841, Queddle Smerby, Campbeltown, Argyllshire, Scotland:$ref[census-1841]
$csvtable{
First,Last,Age,Year
John,McNair,35,1806
Mary,McNair,20,1821
Jean,McNair,1,1840
}
}

partner {
     name Mary Stewart
     gender female
     birth -date 29-may-1816 -location campbeltown -ref mary-birth
     death -date 19-feb-1900 -location campbeltown -ref mary-death
     marriage -date 14-jun-1838 -location campbeltown -ref mary-marriage
     parent stewart-duncan-1781-greenlees
     child {
        name Jane McNair
        gender female
        birth -date 07-jun-1840
        baptism -date 18-jul-1840 -location campbeltown
        death -date 25-dec-1906 -location campbeltown -ref jane1840-death -note umn
     }
     child {
        name Margaret McNair
        gender female
        birth -date 13-mar-1842
        baptism -date 24-apr-1842 -location campbeltown
        death -note dy
     }
     child {
        name Robert McNair
        gender male
        birth -date 16-mar-1844
        baptism -date 02-jun-1844 -location campbeltown
        death -date 19-dec-1927 -location campbeltown -ref robert1844-death -note umn
     }
     child {
        name Anne McNair
        gender female
        birth -date 26-may-1846
        baptism -date 8-jul-1846 -location campbeltown
        death -note young
     }
     child {
        name Duncan Stewart McNair
        gender male
        birth -date 03-jul-1848
        baptism -date 27-aug-1848 -location campbeltown
        death -date 20-jul-1880 -location campbeltown -ref duncan1880-death -note umn
     }
     child {
        name Margaret McNair
        gender female
        birth -date 17-jun-1851
        baptism -date 13-jul-1851 -location campbeltown
        death -date 1857 -location campbeltown -note umn
    }
     child {
        name Sarah McNair
        gender female
        birth -date 02-jul-1853
        baptism -date 01-oct-1853 -location campbeltown
        death -date 10-nov-1921 -location campbeltown -note umn
        burial -name kilchousland -ref sarah1853-burial
     }
     child {
        name Archibald McNair
        gender male
        birth -date 24-dec-1854 -location campbeltown
        death -date 03-jul-1929 -location campbeltown -ref archibald1854-death -note umn
     }
     child {
        name Mary McNair
        gender female
        birth -date 27-jan-1858 -location campbeltown
        death -date 10-feb-1941 -location campbeltown -ref mary1858-death -note umn
     }
}
footnotes {
    census-1841 {
        $elink[https://www.findmypast.com/transcript?id=GBC%2F1841%2F0016597498&expand=true&tab=this]{FindMyPast}
    }
    census-1851 {
        $elink[https://www.findmypast.com/transcript?id=GBC%2F1851%2F0019254600&tab=this]{FindMyPast}    
    }
    census-1861 {
        $elink[https://www.findmypast.com/transcript?id=GBC%2F1861%2F0022167257&expand=true&tab=this]{FindMyPast}
    }
    census-1871 {
        $elink[https://www.findmypast.com/transcript?id=GBC%2F1871%2F0023447125&expand=true&tab=this]{FindMyPast}
    }
    mary-marriage {
        $opr-ref-link[m-1838-507-000-0060-0344 "John McNair" "Mary Stewart"]{
            John McNair farmer Smerby and Mary Stewart
            daughter of Duncan Stewart farmer Peninver 
            were married fourthtenth June 1838
        }
    }
    mary-death {
        $sp-ref-link[d-1900-507-00-0038 0013 "Mary Stewart"]
    }
    jane1840-death {
        $sp-ref-link[d-1908-507-00-0124 0032 "Jeanie MacNair"]
    }
    robert1844-death {
        $sp-ref-link[d-1927-507-00-0104 0034 "Robert MacNair"]
    }
    duncan1848-death {
        $sp-ref-link[d-1880-507-00-0141 0047 "Duncan Stewart McNair"]
    }
    sarah1853-death {
        $sp-ref-link[d-1921-507-00-0091 0031 "Sarah MacNair"]
    } 
    sarah1853-burial {
        https://www.findagrave.com/memorial/159720775/sarah-mcnair
    }
    archibald1854-death {
        $sp-ref-link[d-1929-507-00-0029 0027 "Archibald MacNair"]
    }
    mary1858-death {
        $sp-ref-link[d-1941-507-00-0023  0008 "Mary McNair"]
    }
}
