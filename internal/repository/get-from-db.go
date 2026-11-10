package repository

import (
	"database/sql"
	"pada/internal/model"
)

// Выполняем запрос чтения персональной статистики
func GetPersonalMetrics(db *sql.DB) ([]model.PersonalMetrics, error) {
	query := `SELECT 
		id, fio, dadt, smar, smap, smsd, pct_smar, pct_smap, 
		rvp, tpi, rvpd, tpid, gdp, pss, psa, iis, ipe, l0_4, l1_2 
	FROM personal_metrics`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Обязательно закрываем после работы

	// Инициализируем пустой срез
	var PersonalMetricsList []model.PersonalMetrics

	// Читаем строки по одной
	for rows.Next() {
		var m model.PersonalMetrics

		// Передаем адреса полей структуры в том же порядке, что и в SELECT
		err := rows.Scan(
			&m.ID,
			&m.FIO,
			&m.Dadt,
			&m.Smar,
			&m.Smap,
			&m.Smsd,
			&m.PctSmar,
			&m.PctSmap,
			&m.Rvp,
			&m.Tpi,
			&m.Rvpd,
			&m.Tpid,
			&m.Gdp,
			&m.Pss,
			&m.Psa,
			&m.Iis,
			&m.Ipe,
			&m.L04,
			&m.L12,
		)
		if err != nil {
			return nil, err // Если типы не совпали (например, NULL в базе), вернется ошибка
		}

		// Добавляем запись в срез
		PersonalMetricsList = append(PersonalMetricsList, m)
	}

	// Проверяем, не возникла ли ошибка во время итерации
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return PersonalMetricsList, nil
}

func GetGroupReportList(db *sql.DB) ([]model.GroupReport, error) {
	query := `SELECT id, institute, group_name, starosta, students_list, students_fact, missed_dln, quality_pct, measured_at, starosta_was FROM group_reports`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Обязательно закрываем после работы

	var GroupReportList []model.GroupReport
	for rows.Next() {
		var m model.GroupReport
		err := rows.Scan(
			&m.ID,
			&m.Institute,
			&m.Group_name,
			&m.Starosta,
			&m.Students_list,
			&m.Students_fact,
			&m.Missed_dln,
			&m.Quality_pct,
			&m.Measured_at,
			&m.Starosta_was,
		)

		if err != nil {
			return nil, err // Если типы не совпали (например, NULL в базе), вернется ошибка
		}

		// Добавляем запись в срез
		GroupReportList = append(GroupReportList, m)
	}
	return GroupReportList, nil
}
