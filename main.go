package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// Верхняя панель
	header := tview.NewTextView().
		SetText(" Weatherminal        2024-05-01  12:34   [Moscow]   [S][L][Q] ").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	// Основная погода
	weatherBox := tview.NewTextView().
		SetText("  ☀️  Ясно, +18°C (ощущается как +17°C)\n  Влажность: 45%   Ветер: 3 м/с   Давление: 755 мм рт. ст.").
		SetTextAlign(tview.AlignLeft)

	// Прогноз на 5 дней (таблица)
	forecast := tview.NewTable().
		SetBorders(false)
	days := []string{"Пн", "Вт", "Ср", "Чт", "Пт"}
	icons := []string{"☀️", "🌧", "☁️", "☀️", "🌦"}
	temps := []string{"+18", "+15", "+16", "+19", "+17"}
	for i, day := range days {
		forecast.SetCell(0, i, tview.NewTableCell(day).SetAlign(tview.AlignCenter))
		forecast.SetCell(1, i, tview.NewTableCell(icons[i]).SetAlign(tview.AlignCenter))
		forecast.SetCell(2, i, tview.NewTableCell(temps[i]).SetAlign(tview.AlignCenter))
	}
	forecast.SetFixed(1, 0)

	// Нижняя панель с подсказками
	footer := tview.NewTextView().
		SetText(" [F]orecast  [H]istory  [Fav]orites  [S]ettings  [Q]uit ").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	// Компоновка
	mainFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 1, 0, false).
		AddItem(weatherBox, 3, 0, false).
		AddItem(tview.NewBox().SetBorder(false), 1, 0, false). // Пустая строка
		AddItem(forecast, 5, 0, false).
		AddItem(tview.NewBox().SetBorder(false), 1, 0, false). // Пустая строка
		AddItem(footer, 1, 0, false)

	// Горячие клавиши (моки)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q', 'Q':
			app.Stop()
		case 's', 'S':
			header.SetText(" [Настройки] (мок-экран)  [Q]uit ")
		case 'f', 'F':
			header.SetText(" [Прогноз] (мок-экран)  [Q]uit ")
		case 'h', 'H':
			header.SetText(" [История] (мок-экран)  [Q]uit ")
		case 'l', 'L':
			header.SetText(" [Смена города] (мок-экран)  [Q]uit ")
		}
		return event
	})

	if err := app.SetRoot(mainFlex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
