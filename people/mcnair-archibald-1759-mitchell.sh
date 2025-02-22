name Archibald McNair
gender male
tags Argyll:Campbeltown Immigrant:Canada:Quebec

birth -date 05-mar-1760
death -date 03-aug-1827 -location 'black cape,quebec,canada'
note {
    Maybe immigrated to Quebec in 1821
}
todo {
    An alternate birth date is 9 Nov 1761, however nothing is in the OPR that matches.
}
note {
    Confused with $child-link[mcnair-archibald-1753-mitchell]{Archibald McNair}, b. 1753 who married $i{Jean} Mitchell
}
note {
    Confused with $child-link[mcnair-archibald-1755-bieth]{Archibald McNair}, b. 1755 who married Mary Bieth
}

body {
$p{
Archibald was present on 28-Sep-1824 at the marriage of his son John, according to the records.  In addition, a witness was a "John Jamieson" which was the spouse of his daughter Margaret (who married in 1818).
}
$blockquote{
On this 28th day of September one thousand eight hundred and twenty four, John McNair, farmer, batchelor of major age and Mary McKensie, spinster and minor aged 20 and of New Richmond were married by Banns in the presence and with the consent of his father, relatives and friends and her father, relatives and friends.  The two subscribing witnesses namely,
$br{}
John Jamieson and
$br{}
John Howatson
}
}


body {
Census of 1792, Parish of Kilchousland, Kidonan:
$csvtable{
Age, Year, Men, Women, Children, Role
72, 1720, Achibald McNair,,,Head
61, 1732,                , Margaret Galbreath, , Spouse
27, 1765, Nathan McNair  ,                   , , Son
21, 1771,                , Ann McNair        , , Daughter 
13, 1779,                ,                   , Donald Thomson, TBD
30, 1762, Achibald McNair Junior,            , , Son
11, 1781,                       ,            , Mary McNair, TBD
26, 1766,                       , Margaret Mitchell, , Spouse
 4, 1788,                       ,                  , James McNair, Grand son
 1, 1791,                       ,                  , Margaret McNair, Grand daughter
}
}

external {
    findagrave 220571221
}

partner {
    name Margaret Mitchell
    gender female
    marriage -date 16-aug-1787 -location campbeltown -ref mitchell-marriage

    child {
        name James McNair
        gender male
        baptism -date 02-nov-1788 -location campbeltown
    }
    child {
        name Margaret McNair
        gender female
        baptism -date 25-oct-1790 -location campbeltown -ref margaret1790-baptism
        death -date 15-apr-1884 -location 'New Richmond, Bonaventure, Quebec, Canada'
        partner {
            name John Jamieson
            gender male
            marriage -date 16-jun-1818  -location 'New Richmond, Bonaventure, Quebec, Canada'
        }
    }
    child {
        name Archibald McNair
        gender male
        baptism -date 26-mar-1792 -location campbeltown
    }
    child mcnair-john-1793-mckenzie
}
footnotes {
    mitchell-marriage {
        $opr-ref-link[m-1787-507-000-0020-0170 "Archibald McNair" "Margaret Mitchell"]{
            Archd McNair & Marg Mitchell both of this parish contracted
            16th Aug & were married.
        }
    }
    margaret1790-baptism {
        $opr-ref[b-1790-507-000-0020-0124 "Margaret MacNair"]
    }
}
