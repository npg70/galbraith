name Lachlan Galbreath
birth -date 'about 1778'
death -date sep-1850
burial -name Largnahension Graveyard -location lochgilphead -ref burial
external {
    familysearch M1RG-M1J
    findagrave 160594309
}
note {
    There is another Mary McPhail, born around the same time, who married
    $child-link[galbreath-john-1793-mcphail]{John Galbraith}. Find-a-Grave (in 2023) has them mixed up.
}

tags 'Kilcalmonell and Kilberry'
body {

1841 Census, Baluar, Kilcalmonell, Argyllshire, Scotland:$ref[census1841]:
$csvtable{
First,Last,Age,Year
Lachlan,Galbreith,60,1781
Mary,Galbreith,60,1781
John,Galbreith,30,1811
Margaret,Galbreith,30,1811
Mary,Galbreith,20,1821
Christy,Galbreith,10,1831
Duncan,Galbreith,5,1836
}

}
partner {
    name Mary McPhail
    birth -date 'about 1771'
    death -date 1863
    marriage -date 3-feb-1799 -location 'South Knapdale' -ref mcphail-marriage
    
    child galbraith-donald-1800-brodie


    child galbraith-lachlan-1804-glen

    child {
        name Christine Galbreath
        birth -date 22-jan-1804
        baptism -date 30-jan-1804 -location 'kilcalmonell and kilberry' -ref christine1804-baptism
        death -date 2-dec-1865 -location 'North Bute,Bute' -ref christine1804-death
        partner {
            name John McKeich
        }
    }
    child {
        name Margaret Galbreath
        baptism -date 8-oct-1806 -location 'kilcalmonell and kilberry' -ref margaret1806-baptism
        death -date 1864 -location canada
        partner {
            name Neil Galbreath
            marriage -date 9-aug-1834 -location 'kilcalmonell and kilberry' -ref margaret1806-marriage
        }
        body {
            one son Angus.  They appear in the 1841 and 1851 Census of Campbeltown.  Neil's birth estiamted at 1811.  At age 36, no children, so perhaps none.
        }
    }

    child galbraith-john-1809-mcmillan

    child {
        name Duncan Galbreath
        baptism -date 11-july-1811 -location 'kilcalmonell and kilberry' -ref duncan1811-baptism
        death -date 1839 -location glasgow -ref burial
    }
    child {
        name Mary Galbreath
        baptism -date 6-aug-1818 -location 'south knapdale' -ref mary1818-baptism
        death -date 12-May-1885 -location 'East Williams Township, Middlesex, Ontario, Canada'
        partner {
            name Alexander McLeish
        }
    } 
    child {
        name Isabella Galbraith
        birth -date 'about 1819'
        death -date 13-sep-1878 -location 'Saddell' -ref isabella1819-death
        partner {
            name John McKenzie
            body {
                Gardener in campbeltown
            }
        }
    }
}

footnotes {
    christine1804-baptism {
        $opr-ref-link[b-1804-516-000-0010-0152 "Christine Galbreath"]
    }
    christine1804-death {
        $sp-ref-link[d-1865-557-00-0028 0010 "Christina Galbraith"]
    }
    margaret1806-baptism {
        $opr-ref[b-1806-516-000-0010-0175 "Margaret Galbreath"]
    }
    margaret1806-marriage {
        $opr-ref-link[m-1834-516-000-0020-0388 "Margaret Galbreath" "Niel Galbreath"]
    }
    duncan1811-baptism {
        $opr-ref[b-1811-516-000-0010-0212 "Duncan Galbreath"]
    }
    mary1818-baptism {
        $opr-ref-link[b-1818-533-000-0010-0124 "Mary Galbreath"]
    }
    isabella1819-death {
        $sp-ref-link[d-1878-531-01-0007 0003 "Isabella Galbraith"]
    }
    mcphail-marriage {
        $opr-ref-link[m-1799-533-000-0010-0159 "Lachline Galbreath" "Mary McPhail"]{
        3 Feb Registration for Marriage Lachline Galbreath and Mary Mc Phail
        both in this parish
        }
    }
    census1841 {
        https://www.findmypast.com/transcript?id=GBC/1841/0016609067&expand=true
    }
    burial {
    Original transcription of the tombstone for Lachlan Galbraith, as shown on $elink[https://www.findagrave.com/memorial/160594309/lachlan-galbraith]{Find-A-Grave}:
        $blockquote{$pre{
Erected
by
Mary McPheil
In Memory of her Husband
Lachlan Galbreath
Tenant Balure Who Died
? Sept 1850 Aged 72
Also
Their Son Duncan who died
? Dec 1836 Age 27}}

    }
}
