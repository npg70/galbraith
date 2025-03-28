name John Galbreath
gender male
baptism -date 24-jan-1815 -location 'killean and kilchenzie' -ref john1815-baptism
death -date 15-Feb-1890 -location campbeltown -ref death
external {
    familysearch LZ62-VM4
}
tags Argyll:Campbeltown
body {
    He was a distillery workman.
}

todo census { needs fixing }

body {
Census of 1851, Burn Side Street, Campbeltown, Cantyre, Argyllshire, Scotland:$ref[census1851]
$csvtable{

}
}

body {
Census of 1861, Ladymary Row, Campbeltown, Argyllshire, Scotland:$ref[census1861]
$csvtable{
First,Last,Role,Status,Age,Year,Birth Place,Job
John,Galbraith,Head,Married,47,1814,Kilcalmonell,Distillery man
Mary,Galbraith,Wife,Married,45,1816,Campbeltown,Wife of
Archibald,Galbraith,Son,-,18,1843,Campbeltown,Malster
John,Galbraith,Son,-,13,1848,Campbeltown,Tailor
Duncan,Galbraith,Son,-,10,1851,Campbeltown,Scholar
Annabella,Beith,Mother in law,Widow,70,1791,Killean,x-Housekeep
}
}
body {
Census of 1871, Saddell Street, Campbeltown, Argyllshire, Scotland:$ref[census1871]
$csvtable{
First,Last,Role,Status,Age,Year,Place,Job
John,Galbreath,Head,Married,57,1814,Clachan,Laborer
Mary,Galbreath,Wife,Married,56,1815,"Kilkevan",-
}
}
partner {
    name Mary McTaggart
    gender female
    baptism -date 01-sep-1812 -location campbeltown
    death -date 14-feb-1874 -location campbeltown
    parent {
        name Alexander McTaggart
        partner {
            name Annabella Beith
        }
    }
    marriage -date 27-Nov-1839 -location campbeltown -ref marriage
    child {
        name Annabella Galbreath
        baptism -date 29-nov-1840 -location campbeltown -ref annabella1840-baptism
    }
    child galbreath-archibald-1842-omay
    child galbreath-alexander-1844-campbell-matheson 
    child {
        name John Galbreath
        baptism -date 8-dec-1845 -location campbeltown -ref john1845-baptism
        death -date 5-jun-1916 -location greenock -ref john1845-death -note umd
    }
    child {
        name Duncan Galbraith
        baptism -date 12-feb-1851 -location campbeltown -ref duncan1851-baptism
        death -date 10-mar-1885 -location southend -ref duncan1851-death -note umd
    }
}

footnotes {
    john1815-baptism {
        $opr-ref-link[b-1815-519-000-0010-0277 "John Galbreath"]{
            July | 24 | John, Lawful Son to John Galbreath and Margt
            McGill Innkeeper Tayinloan[?]
        }
    }
    death {
        $sp-ref-link[d-1890-507-00-0026 0009 "John Galbraith"]
    }
    marriage {
        $opr-ref-link[m-1839-507-000-0060-0358 "John Galbrath" "Mary McTaggart"]{
            John Galbrath laborer Machrehanish 
            and Mary McTaggart  Daughter of
            Alexander McTaggart laborer $i{illegible}
            both of this Parish were married
            twenty seventh November 1839.
        }
    }
    annabella1840-baptism {
        $opr-ref[b-1840-507-000-0070-0184 "Annabella Galbreath"]
    }
    john1845-baptism {
        $opr-ref[b-1847-507-000-0070-0328 "John Galbreath"]
    }
    john1845-death {
        $sp-ref-link[d-1916-564-02-0350 0117 "John Galbraith"]    
    }
    duncan1851-baptism {
        $opr-ref[b-1851-507-000-0070-0380 "Duncan Galbraith"]
    }
    duncan1851-death {
        $sp-ref-link[d-1885-532-00-0004 0002 "Duncan Galbraith"]
    }
    census1851 {
        https://www.findmypast.com/transcript?id=GBC/1851/0019256008&expand=true
    }
    census1861 {
        https://www.findmypast.com/transcript?id=GBC/1861/0022162723&expand=true
    }
    census1871 {
        https://www.findmypast.com/transcript?id=GBC%2F1871%2F0023441340&tab=this    
    }
    census1881 {
        https://www.findmypast.com/transcript?id=GBC%2F1881%2F0029344266&tab=this
    }
}
