name Donald Galbraith
baptism -date 1800 -location 'kilcalmonell and kilberry' -ref donald1800-baptism
death -date 1839 -ref donald1800-death
tags 'Kilcalmonell and Kilberry'
body {

Mary in 1841 Census Teritigan, Kilcalmonell, Argyllshire, Scotland:$ref[census1841]
$csvtable{
First,Last,Age,Year
Mary,Galbraith,40,1801
Lachland,Galbraith,18,1823
Mary,Galbraith,16,1825
Donald,Galbraith,6,1835
John,Galbraith,9,1832
}
Note that the 1841 Census often rounded ages to nearest multiple of 5.

1851 Census, Tiretigan, Kilcalmonell, Kintyre, Argyllshire, Scotland: $ref[census1851]
$csvtable{
First,Last,Role,Age,Year,Birth
Mary,McBride,Widow,50,1801,South Knapdale
Lachlan,Galbraith,Son,24,1827,South Knapdale
Duncan,Galbraith,Son,14,1837,Kilberry
Donald,Galbraith,Son,12,1839,Kilberry
Neil,McBride,Visitor,32,1819,South Knapdale
}


1861 Census, East Cretshengan, Kilberry, Argyllshire, Scotland: $ref[census1861]
$csvtable{
First,Last,Role,Status,Age,Year,Place,Occupation
Donald,Galbraith,Head,Unmarried,22,1839,Kilberry,Fisherman
Mary,Galbraith,Widow,Widow,60,1801,"ORMSARAY",Farm Laborer
}

1871 Census, Tarbert, Argyllshire, Scotland:$ref[census1871]
$csvtable{
First,Last,Role,Status,Age,Year,Place,Occupation
Mary,Galbraith,Head,Widow,71,1800,South Knapdale,Former Domestic Servant
Donald,Galbraith,Son,Unmarried,30,1841,Kilberry,Fisherman
}
}
partner {
    name Mary Brodie/McBride
    birth -date 'about 1800'
    marriage -date 24-dec-1825 -location 'Kilcalmonell and Kilberry' -ref brodie-marriage
    death -date 20-jan-1885 -location Tarbert -ref brodie-death

    child {
        name Lachlan Galbreath
        baptism -date 3-jan-1827 -location 'South Knapdale' -ref lachlan1827-baptism
        death -date 22-nov-1888 -location greenock -ref lachlan1827-death
        partner {
            name Jane Templeton
        }
    }
    child {
        name Mary Galbreath
        baptism -date 15-jan-1829 -location 'South Knapdale' -ref mary1829-baptism
    }
    child {
        name Christian Galbraith
        baptism -date 31-jul-1830 -location 'Kilcalmonell and Kilberry' -ref christian1830-baptism
    }
    child {
        name John Galbraith
        baptism -date 15-jul-1832 -location 'Kilcalmonell and Kilberry' -ref john1832-baptism
    }
    child {
        name Isabella Galbraith
        baptism -date 6-july-1834 -location 'Kilcalmonell and Kilberry' -ref isabella1834-baptism
    }
    child galbraith-duncan-1836-scott

    child {
        name Donald Galbraith
        birth -date 'about 1839'
        death -date 8-mar-1912 -location tarbert -ref donald1839-death
        partner {
            name Grace Stalker
            marriage -date 29-jan-1874 -location skipness -ref donald1839-marriage
        }
    }
}

footnotes { 
    census1841 {
        https://www.findmypast.com/transcript?id=GBC/1841/0016620751&expand=true
    }
    census1851 {
        https://www.findmypast.com/transcript?id=GBC/1851/0019282882
    }
    census1861 {
        https://www.findmypast.com/transcript?id=GBC%2F1861%2F0022460225
    }
    census1871 {
        https://www.findmypast.com/transcript?id=GBC%2F1871%2F0023495931
    }   
    donald1800-baptism {
        $opr-ref-link[b-1800-516-000-0010-0121 "Donald Galbreath"]
        Only marked as the year 1800.  Scotlands People for some reason lists the unknown date to be "31 Dec 1799"
    }
    brodie-marriage {
        $opr-ref-link[m-1825-516-000-0020-0372 "Donald Galbraith" "Mary Brodie"]{
            The which day were booked in order to marriage
            Donald Galbraith Balure + Mary Brodie in Barinlongart parish of
            South Knapdale and proclaimed the 25th Dec 1825
        }
    }   
    brodie-death {
        $sp-ref-link[d-1885-535-00-0003 0001 "Mary Galbraith"]{
            name: Mary Galbraith, Widower of Donald Galbraith, Farmer;
            date: 1885 January Twentieth, Coalside Cottages, Tarbert, South Knapdale;
            Age: 85 $i{birth 1800};
            Father: Dugald McBride, Farmer (deceased)
            Mother: Mary Morrison (deceased);
            Present: Donald Galbraith, Son
        }
    }
    lachlan1827-baptism {
        $opr-ref[b-1827-533-000-0020-0038 "Lachline Galbreath"]. Mother listed as "Mary Brodie."
    }
    lachlan1827-death {
        $sp-ref-link[d-1888-564-03-0616 0206 "Lachlan Galbreath"]. Mother listed as "McBride."
    }
    mary1829-baptism {
        $opr-ref[b-1829-533-000-0020-0051 "Mary Galbreath"]. Mother listed as "Mary Brody."
    }
    christian1830-baptism {
        $opr-ref[b-1830-516-000-0020-0319 "Christian Galbraith"]. Mother listed as "Mary Brodie."
    }
    john1832-baptism {
        $opr-ref[b-1832-516-000-0020-0330 "John Galbraith"]. Mother listed as "Mary Bride."
    }
    isabella1834-baptism {
        $opr-ref[b-1834-516-000-0020-0337 "Isabella Galbraith"]. Mother listed as "Mary Brody."
    }

donald1800-death {
His death does not appear to be recorded, however from $elink[https://archive.org/details/CommemorativeBiographicalRecordOfTheCountyOfKentOntario]{Commemorative Biographical Record Of The County Of Kent Ontario}, published in 1904, by J.H. Beers & Co, page 278 under the "Duncan Galbraith" section has this about Donald, the first child of Lachlan Galbraitha and Mary McPhail:
 $blockquote{
 Donald, born in 1800, married in Scotland, made several voyages to Canada, and died in 1839. 
}
}
    donald1839-death {
        $sp-ref-link[d-1912-535-00-0006 0002 "Donald Galbraith"].  Married but spouse not listed.
    }
    donald1839-marriage {
        $sp-ref-link[m-1874-531-02-0001 0001 "Donald Galbraith" "Grace Stalker"]. Mother listed as "McBride"
    }
}
