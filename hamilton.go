package main

var aboutHamilton string = "Шкала тревоги Гамильтона (англ. The Hamilton Anxiety Rating Scale, сокр. HARS) - клиническая рейтинговая шкала, предназначенная для измерения тяжести тревожных расстройств. Шкала состоит из 14 пунктов. По каждому пункту необходимо выбрать значение, наиболее соответствующее степени выраженности симптомов."

var questionsHamilton [14]string = [14]string{
	"1. Тревожное настроение (озабоченность, ожидание наихудшего, тревожные опасения, раздражительность)",
	"2. Напряжение (ощущение напряжения, вздрагивание, легко возникающая плаксивость, дрожь, чувство беспокойства, неспособность расслабиться)",
	"3. Страхи (темноты, незнакомцев, одиночества, животных, толпы, транспорта)",
	"4. Инсомния (затрудненное засыпание, прерывистый сон, не приносящий отдыха, чувство разбитости и слабости при пробуждении, кошмарные сны)",
	"5. Интеллектуальные нарушения (затруднение концентрации внимания, ухудшение памяти)",
	"6. Депрессивное настроение (утрата привычных интересов, чувства удовольствия от хобби, подавленность, ранние пробуждения, суточные колебания настроения)",
	"7. Соматические мышечные симптомы (боли, подергивания, напряжение, судороги клонические, скрипение зубами, срывающийся голос, повышенный мышечный тонус)",
	"8. Соматические сенсорные симптомы (звон в ушах, нечеткость зрения, приливы жара и холода, ощущение слабости, покалывания)",
	"9. Сердечно-сосудистые симптомы (тахикардия, сердцебиение, боль в груди, пульсация в сосудах, частые вздохи)",
	"10. Респираторные симптомы (давление и сжатие в груди, удушье, частые вздохи)",
	"11. Гастроинтестинальные симптомы (затрудненное глотание, метеоризм, боль в животе, изжога, чувство переполненного желудка, тошнота, рвота, урчание в животе, диарея, запоры, снижение веса тела)",
	"12. Мочеполовые симптомы (учащенное мочеиспускание, сильные позывы на мочеиспускание, аменорея, менорагия, фригидность, преждевременная эякуляция, утрата либидо, импотенция)",
	"13. Вегетативные симптомы (сухость во рту, покраснение или бледность кожи, потливость, головные боли с чувством напряжения)",
	"14. Поведение при осмотре (ёрзанье на стуле, беспокойная жестикуляция и походка, тремор, нахмуривание лица, напряженное выражение лица, вздохи или учащенное дыхание, частое сглатывание слюны)",
}

var point = map[string]int{
	"В слабой степени":        1,
	"В умеренной степени":     2,
	"В тяжёлой степени":       3,
	"В очень тяжёлой степени": 4,
}

var resultHamilton = map[int]string{
	17: "Отсутствие тревоги.",
	24: "Средняя выраженность тревоги.",
	25: "Сильная тревога.",
}
