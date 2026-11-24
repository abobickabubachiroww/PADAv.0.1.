import {
	GetPersonalMetricsByGroup,
	SearchGroupReports,
} from '../wailsjs/go/ui/WailsApp'

const contentContainer = document.getElementById('app-content')

const routesMap = {
	'group-report': 'groups',
	'personal-metrics': 'personalMetrics',
}

// ---------- personal metrics ----------

async function loadTable() {
	const tbody = document.getElementById('table-body')
	if (!tbody) return

	const groupFilter = window.__pendingGroupFilter || ''
	window.__pendingGroupFilter = null

	const title = document.querySelector('.page h1')
	if (title) {
		title.textContent = groupFilter
			? `Личные показатели — группа ${groupFilter}`
			: 'Личные показатели'
	}

	if (!groupFilter) {
		tbody.innerHTML =
			'<tr><td colspan="17">Выберите группу на главной странице</td></tr>'
		return
	}

	tbody.innerHTML = '<tr><td colspan="17">Загрузка данных...</td></tr>'

	try {
		const data = await GetPersonalMetricsByGroup(groupFilter)
		tbody.innerHTML = ''

		if (!data || data.length === 0) {
			tbody.innerHTML = `<tr><td colspan="17">В группе «${groupFilter}» данных нет</td></tr>`
			return
		}
		const medals = ['🥇', '🥈', '🥉']
		const topTPI = new Map(
			data
				.filter(
					item =>
						item.FIO !== 'СРЕДНИЕ ПОКАЗАТЕЛИ' &&
						item.TPI != null &&
						!isNaN(Number(item.TPI)),
				)
				.sort((a, b) => Number(b.TPI) - Number(a.TPI))
				.slice(0, 3)
				.map((item, i) => [item, medals[i]]),
		)
		data.forEach(item => {
			const tr = document.createElement('tr')
			const safeNum = val => (val != null ? Number(val).toFixed(2) : '-')

			const fields = [
				item.ID,
				item.FIO,
				item.Dadt,
				safeNum(item.CnStr),
				safeNum(item.RVP),
				safeNum(item.TPI),
				safeNum(item.RVPD),
				safeNum(item.TPID),
				safeNum(item.GDP),
				safeNum(item.PSS),
				safeNum(item.PSA),
				safeNum(item.IIS),
				safeNum(item.IPE),
				safeNum(item.L0),
				safeNum(item.L1),
				safeNum(item.L2),
				safeNum(item.L3),
			]

			fields.forEach((val, idx) => {
				const td = document.createElement('td')
				td.textContent = val

				// Помечаем критические значени Диазадаптации/RVP/TPI
				if (
					idx === 2 &&
					(item.Dadt === 'критическая' || item.Dadt === 'сильная')
				) {
					td.classList.add('critical-cell')
				}
				if (idx === 4 && item.RVP <= 0) {
					td.classList.add('critical-cell')
				}
				if (idx === 5 && item.TPI <= 0) {
					td.classList.add('critical-cell')
				}

				if (idx === 1 && topTPI.has(item)) {
					td.classList.add('top-medal-cell')
					td.textContent = `${topTPI.get(item)} ${val}`
				}

				tr.appendChild(td)
			})

			// СНАРУЖИ цикла — работаем с tr
			if (item.FIO === 'СРЕДНИЕ ПОКАЗАТЕЛИ') {
				tr.classList.add('averages-row')
			}

			tbody.appendChild(tr)
		})
	} catch (err) {
		console.error('Ошибка:', err)
		tbody.innerHTML = `<tr><td colspan="17" style="color: red;">Ошибка: ${err}</td></tr>`
	}
}

// ---------- group reports ----------
function renderGroupRows(items) {
	const tbody = document.getElementById('group-table-body')
	if (!tbody) return
	tbody.innerHTML = ''

	if (!items || items.length === 0) {
		tbody.innerHTML = '<tr><td colspan="10">Ничего не найдено</td></tr>'
		return
	}

	items.forEach(item => {
		const tr = document.createElement('tr')

		const dateStr = item.Measured_at
			? new Date(item.Measured_at).toLocaleDateString('ru-RU')
			: '-'
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

		fields.forEach((val, idx) => {
			const td = document.createElement('td')
			td.textContent =
				val !== undefined && val !== null && val !== '' ? val : '-'

			if (idx === 2 && item.Group_name) {
				td.classList.add('group-link')
				td.dataset.group = item.Group_name
			}

			tr.appendChild(td)
		})
		tbody.appendChild(tr)
	})
}

function showEmptyGroupTable() {
	const tbody = document.getElementById('group-table-body')
	if (!tbody) return
	tbody.innerHTML =
		'<tr><td colspan="10">Введите институт, год и полугодие, затем нажмите «Поиск»</td></tr>'
}

// ---------- routing ----------

async function loadPage(route) {
	const filename = routesMap[route] || 'groups'

	try {
		const response = await fetch(`/pages/${filename}.html`)
		if (!response.ok) throw new Error(`Страница "${filename}.html" не найдена`)

		const html = await response.text()
		contentContainer.innerHTML = html

		if (route === 'personal-metrics') {
			loadTable()
		} else if (route === 'group-report') {
			showEmptyGroupTable()
		}
	} catch (error) {
		console.error('Ошибка загрузки страницы:', error)
		contentContainer.innerHTML = `<h1 style="color:red">404</h1><p>${error.message}</p>`
	}
}

window.addEventListener('hashchange', () => {
	const route = window.location.hash.replace('#/', '') || 'group-report'
	loadPage(route)
})

document.addEventListener('DOMContentLoaded', () => {
	const initialRoute = window.location.hash.replace('#/', '') || 'group-report'
	loadPage(initialRoute)
})

// ---------- поиск ----------

document.body.addEventListener('submit', async e => {
	if (e.target.id !== 'group-search-form') return
	e.preventDefault()

	const fd = new FormData(e.target)
	const payload = {
		institute: String(fd.get('institute') || '').trim(),
		year: Number(fd.get('year')),
		half: String(fd.get('half') || '').trim(),
	}

	const tbody = document.getElementById('group-table-body')
	if (tbody) tbody.innerHTML = '<tr><td colspan="10">Поиск...</td></tr>'

	try {
		const rows = await SearchGroupReports(payload)
		renderGroupRows(rows)
	} catch (err) {
		console.error('Ошибка поиска:', err)
		if (tbody) {
			tbody.innerHTML = `<tr><td colspan="10" style="color:red">Ошибка: ${err}</td></tr>`
		}
	}
})

// ---------- клики: группа + кнопка «назад» ----------

document.body.addEventListener('click', e => {
	// «назад» — на главную страницу
	if (e.target.closest('.back-btn')) {
		window.location.hash = '#/group-report'
		return
	}

	// клик по названию группы
	const cell = e.target.closest('.group-link')
	if (!cell) return

	const groupName = cell.dataset.group
	if (!groupName) return

	window.__pendingGroupFilter = groupName
	window.location.hash = '#/personal-metrics'
})
