// 1. ИМПОРТИРУЕМ ОБЕ ФУНКЦИИ
import {
	GetPersonalMetrics,
	GetGroupReportList,
} from '../wailsjs/go/ui/WailsApp'

const contentContainer = document.getElementById('app-content')

const routesMap = {
	'group-report': 'groups',
	'personal-metrics': 'personalMetrics',
}

// ==========================================
// 1. ЛОГИКА ДЛЯ ЛИЧНЫХ ПОКАЗАТЕЛЕЙ
// ==========================================
async function loadTable() {
	const tbody = document.getElementById('table-body')
	if (!tbody) return

	tbody.innerHTML = '<tr><td colspan="18">Загрузка данных из Go...</td></tr>'

	try {
		const data = await GetPersonalMetrics()
		tbody.innerHTML = ''

		if (!data || data.length === 0) {
			tbody.innerHTML = '<tr><td colspan="18">Данных нет</td></tr>'
			return
		}

		data.forEach(item => {
			const tr = document.createElement('tr')
			const safeNum = val => (val != null ? Number(val).toFixed(2) : '-')

			const fields = [
				item.ID,
				item.FIO,
				item.Dadt,
				safeNum(item.Smar),
				safeNum(item.Smap),
				safeNum(item.Smsd),
				safeNum(item.PctSmar),
				safeNum(item.PctSmap),
				safeNum(item.Rvp),
				safeNum(item.Tpi),
				safeNum(item.Rvpd),
				safeNum(item.Tpid),
				safeNum(item.Gdp),
				safeNum(item.Pss),
				safeNum(item.Psa),
				safeNum(item.Iis),
				safeNum(item.Ipe),
				safeNum(item.L04),
				safeNum(item.L12),
			]

			fields.forEach(val => {
				const td = document.createElement('td')
				td.textContent = val
				tr.appendChild(td)
			})
			tbody.appendChild(tr)
		})
	} catch (err) {
		console.error('Ошибка:', err)
		tbody.innerHTML = `<tr><td colspan="19" style="color: red;">Ошибка: ${err}</td></tr>`
	}
}

// ==========================================
// 2. ЛОГИКА ДЛЯ ГРУППОВОГО ОТЧЕТА (НОВАЯ)
// ==========================================
async function loadGroupTable() {
	const tbody = document.getElementById('group-table-body')
	if (!tbody) return

	tbody.innerHTML = '<tr><td colspan="10">Загрузка данных из Go...</td></tr>'

	try {
		const data = await GetGroupReportList()
		tbody.innerHTML = ''

		if (!data || data.length === 0) {
			tbody.innerHTML = '<tr><td colspan="10">Данных нет</td></tr>'
			return
		}

		data.forEach(item => {
			const tr = document.createElement('tr')

			// Форматируем дату (обрезаем время)
			const dateStr = item.Measured_at
				? new Date(item.Measured_at).toLocaleDateString('ru-RU')
				: '-'
			// Форматируем boolean в текст
			const qualityStr = item.Quality_pct != null ? item.Quality_pct + '%' : '-'

			const fields = [
				item.ID,
				item.Institute,
				item.Group_name,
				item.Starosta,
				item.Students_list,
				item.Students_fact,
				item.Missed_dln,
				qualityStr,
				dateStr,
				item.Starosta_was,
			]

			fields.forEach(val => {
				const td = document.createElement('td')
				td.textContent =
					val !== undefined && val !== null && val !== '' ? val : '-'
				tr.appendChild(td)
			})
			tbody.appendChild(tr)
		})
	} catch (err) {
		console.error('Ошибка:', err)
		tbody.innerHTML = `<tr><td colspan="10" style="color: red;">Ошибка: ${err}</td></tr>`
	}
}

// ==========================================
// 3. РОУТИНГ
// ==========================================
async function loadPage(route) {
	const filename = routesMap[route] || 'groups'

	try {
		const response = await fetch(`/pages/${filename}.html`)
		if (!response.ok) throw new Error(`Файл "${filename}.html" не найден`)

		const html = await response.text()
		contentContainer.innerHTML = html

		// Запускаем нужную функцию в зависимости от страницы
		if (route === 'personal-metrics') {
			loadTable()
		} else if (route === 'group-report') {
			loadGroupTable()
		}
	} catch (error) {
		console.error('Ошибка загрузки страницы:', error)
		contentContainer.innerHTML = `<h1 style="color:red">404</h1><p>${error.message}</p>`
	}
}

// ДЕЛЕГИРОВАНИЕ СОБЫТИЙ
document.body.addEventListener('click', e => {
	const btn = e.target.closest('.nav-btn')
	if (btn && btn.dataset.route) {
		window.location.hash = `#/${btn.dataset.route}`
	}
})

window.addEventListener('hashchange', () => {
	const route = window.location.hash.replace('#/', '') || 'group-report'
	loadPage(route)
})

document.addEventListener('DOMContentLoaded', () => {
	const initialRoute = window.location.hash.replace('#/', '') || 'group-report'
	loadPage(initialRoute)
})
