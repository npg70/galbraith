name Mary Galbreath	
tags Argyll:Campbeltown
baptism -date 28-Apr-1751 -location Campbeltown -ref mary1751-baptism
death -date 03-apr-1825 -location campbeltown -ref mary1751-death
todo {
    link to kilkerran
}
note {
    There is another Mary Galbreath born in 1751, the daughter of $child-link[galbreath-david-1719-langwill]{DAvid Galbreath} and Jean Langwill.
}

body {
Census of 1792, 53 Kirk Street:

$csvtable{
Age,Year,Men,Women,Children
40,1752,James Park,,
38,1754,,Mary Langwill,
14,1778,,,John Park
 8,1784,,,Mary Park
 6,1786,,,James Park
 4,1788,,,Betty Park
19,1773,,Margaret Muir,
32,1760,William Greenlees,,
19,1773,,Agnes Andrew,
52,1740,,Mary Ferguson,
}
}

partner {
    name Matthew Andrew
    marriage -date 15-nov-1771 -location campbeltown -ref andrew-marriage

    child andrew-agnes-1772-greenlees

    child {
        name Elisabeth Andrew
        baptism -date 18-jan-1775 -location campbeltown
    }
}

partner {
    name James Park
    birth -date 'about 1750'
    death -date 09-oct-1806 -location campbeltown

    marriage -date 'about 1776'

    child {
        name James/John Park
        baptism -date 10-jul-1777 -location campbeltown
    }
    child {
        name Robert Park
        baptism -date 04-jul-1779 -location campbeltown
        body {
            Died young.
        }
    }

    child park-mary-1782-mclachlan

    child {
        name James Park 
        baptism -date 29-aug-1784 -location campbeltown
        death -date 1793
    }
    child park-elizabeth-1786-reid 
    child {
        name Margaret Park
        birth -date 02-oct-1790
        baptism -date 04-oct-1790 -location campbeltown -ref marg1790-baptism
        death -date 1791
    }
    child {
        name Margaret Park
        birth -date 24-aug-1794
        baptism -date 26-aug-1794 -location campbeltown -ref marg1794-baptism
    }
}

footnotes {
    mary1751-baptism {
        $opr-ref-link[b-1751-507-000-0011-0276 "Mary Galbreath"]{
            Mary | John Galbreath and Agnes Langwill had a Dau
            baptized 28th April named Mary
        }
    }

    mary1751-death {
       From Kilkerran Cemetery:
$blockquote{
1815 Erected by DONALD McLACHLAN, Malster in Campbeltown in memory of his father in law J. PARK late merchant there who died 9 October 1806 aged 56 years also his mother in law MARY GALBREATH who died 3 April 1825 aged 71 years.}
    }
    andrew-marriage {
        $opr-ref-link[m-1771-507-000-0011-0473 "Mary Galbreath" "Matthew Andrew"]{
            Matthew Andrew and Mary Galbreath in this Parish - Nov 15th
        }
        $opr-ref-link[m-1771-507-000-0020-0156 "Mary Galbreath" "Matthew Andrew"]{
            Nov 27th Matthew Andrew & Mary Galbreath in this Parish both
        }
        $opr-ref-link[m-1771-507-000-0010-0343 "Mary Galbreathson" "Matthew Andrew"]{
        Nov 1th Matthew Andrew and Mary Galbreath in this parish
    }
    }
    marg1790-baptism {
        $opr-ref-link[b-1790-507-000-0020-0123 "Margaret Park"]{
          Margaret Lawful dau. to James Park Mercht in Town
          & Mary Galbreath was born 2nd & Bapt 4th Oct. 1790
        }
    }
    marg1794-baptism {
        $opr-ref-link[b-1794-507-000-0040-0118 "Margaret Park"]{
            $i{August 1794} Margaret Lawful Dau to James Park & Mary Galbreath
            in Town born 24th Bapt 26th
        }
    }
}
