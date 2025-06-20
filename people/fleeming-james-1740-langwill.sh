name James Fleeming
gender male
tags Argyll:Campbeltown
external {
    familysearch 9M9D-QL7
}
birth -date 'say 1740'

confused-with fleeming-john-1703-kirkland
confused-with fleeming-john-1713-langwill

body {
$ul{
$li{"Written 22 Jan 1785"}
$li{"his spouse, Janet Langwill", thus she was alive in 1785}
$li{"... his only son, John Fleeming ..."}
$li{"He debarred the rest of his children apart from... (the following"}
$li{"Isobel, the eldest, in America"}
$li{"Janet, spouse of John McMillan, shipmaster in Campbeltown"}
$li{"Margaret, spouse of Andrew Ralston in Knockstaple"}
$li{"Agnes, spouse of Alexander Galbraith, tailor in Campbeltown"}
$li{"Mary, spouse of John Paterson, cooper in Campbeltown"}
$li{"Jean youngest daughter .. until she reached the age 18"}
}

$h3{What we know}

$ul{
$li{James Fleeming married Janet Langwill in 1761.}
$li{Janet was alive in 1785}
$li{Child: Agnes, Mary, Jean all match}
$li{$child-link[fleeming-john-1703-kirkland]{John Fleeming} who married Jean Kirkland is someone else.}
$li{Jean is in the census of 1792, living with John Fleeming and Jean Colville, of age 12, which matched everything}
}

$h3{Very Likely}

$ul{
$li{Child Margaret was born around 1765, married Andrew Ralston.  Birth appears not recorded, but age is based on census and gravestones}
$li{Janet is not in the 1792 Census, implying she died before 1792}
$li{John was actually "James" and/or born in 1768-1769:
  	$ul{

$li{John Fleeming with Jean Colville is in the census of 1792. He is listed as 21, which gives a birth of 1771.  They are off two years off, so birth is probably 1769.  Jean Colvills listd as 19.
}
$li{John's burial on find-a-grave, gives a birth of 1768}
}
}
}

$h3{Still unknown}

$ul{
$li{ Child John, Isobel and Janet have missed records, or indicate birth before 1761.
}
}

}
partner {
    name Janet Langwill
    gender female
    marriage -date 09-jun-1761 -location campbeltown -ref langwill-marriage


    child fleeming-janet-1761-mcmillan

    child {
        name Agnes Fleeming
        gender female
	birth -date 29-oct-1762
        baptism -date 31-oct-1762 -location campbeltown -ref agnes1762-baptism
        partner galbreath-alexander-1753-flemming
    }

    child {
        name Margaret Fleeming
        gender female
        birth -date 'about 1765'
        partner ralston-andrew-1758-fleeming
    }

    child fleeming-mary-1765-paterson

    child {
	name James Fleeming
	gender male
	baptism -date 17-oct-1769 -location campbeltown -ref james1769-baptism
	body {
		TBD
	}
    }

    child {
	name Elizabeth Fleeming
	gender female
	birth -date 21-mar-1775
	baptism -date 23-mar-1775 -location campbeltown -ref eliza1775-baptism
	body {
		assumed died young.
	}
    }

    child fleeming-jean-1778-mcmurchy


}
footnotes {
    langwill-marriage {
        $opr-ref-link[m-1761-507-000-0011-0461 "John Fleming" "Janet Langville"]
        $blockquote{
            James Fleming and Janet Langville both of this Parish June 9th [1761]
        }
    }
    agnes1762-baptism {
	$opr-ref-link[b-1762-507-000-0011-0332 "Agnes Fleeming"]
	$blockquote{
		Agnes | James Fleeming and Janet Langwill had a Daughter
		born the 29th and baptized 31st Oct named Agnes.
	}
	$opr-ref[b-1762-507-000-0010-0271 "Angus Fliming"]
    }
    mary1765-baptism {
	$opr-ref-link[b-1765-507-000-0011-0348 "Mary Fleeming"]
	$blockquote{
		Mary | James Fleeming and Janet Langwill had a
		Dau. born 9th and baptized 10th Oct named Mary
	}
	$opr-ref-link[b-1765-507-000-0010-0296 "Mary Fleeming"]

    }
    james1769-baptism {
        $opr-ref-link[b-1769-507-000-0011-0372 "James Fleeming"]
	$blockquote{
		James | James Fleeming and Janat Langwill had a son
		born the 17th Oct named James
        }
	$opr-ref[b-1769-507-000-0010-0336 "James Fleeming"]
    }
    eliza1775-baptism {
	$opr-ref-link[b-1775-507-000-0020-0017 "Elisabeth Fleeming"]
	$blockquote{
		Elisabeth | Daughter to James Fleeming and Jenat Longwill
		Born 21st March Bap 23rd named Elisabeth	
	}
    }
}
