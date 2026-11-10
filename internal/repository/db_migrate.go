package repository

import (
	"database/sql"
	"fmt"
)

// развертывание БД, если ее нет
func Migrate(db *sql.DB) error {
	fmt.Println("DB created")
	schema := ` 
	DROP TABLE IF EXISTS personal_metrics;
	DROP TABLE IF EXISTS group_reports;
	CREATE TABLE IF NOT EXISTS personal_metrics (
		id         INTEGER PRIMARY KEY AUTOINCREMENT,
		fio        TEXT NOT NULL,
		dadt			 text,
		smar       REAL,
		smap       REAL,
		smsd       REAL,
		pct_smar   REAL,   -- %SMAR
		pct_smap   REAL,   -- %SMAP
		rvp        REAL,
		tpi        REAL,
		rvpd       REAL,
		tpid       REAL,
		gdp        REAL,
		pss        REAL,
		psa        REAL,
		iis        REAL,
		ipe        REAL,
		l0_4       REAL,
		l1_2       REAL,
		l2_1_2     REAL,   -- L2_1,2 
		l3_3       REAL,
		special	   int, -- сво овз звезды

		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS group_reports (
		id             INTEGER PRIMARY KEY AUTOINCREMENT,
		institute			 TEXT    NOT NULL,
		group_name     TEXT    NOT NULL,
		starosta       TEXT    NOT NULL,
		students_list  INTEGER NOT NULL,
		students_fact  INTEGER NOT NULL,
		missed_dln		 INTEGER NOT NULL,
		quality_pct    REAL NOT NULL,
		measured_at    DATE    NOT NULL,
		starosta_was   TEXT NOT NULL ,

		created_at     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_group_reports_group ON group_reports(group_name);
	CREATE INDEX IF NOT EXISTS idx_group_reports_date  ON group_reports(measured_at);
	CREATE INDEX IF NOT EXISTS idx_metrics_fio ON personal_metrics(fio);
	`
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	// Тестовые данные. INSERT ... SELECT ... WHERE NOT EXISTS делает вставку
	// идемпотентной: повторный вызов Migrate не создаст дубликатов.
	testRows := `
		INSERT INTO personal_metrics (fio, dadt, smar, smap, smsd, pct_smar, pct_smap, rvp, tpi, rvpd, tpid, gdp, pss, psa, iis, ipe, l0_4, l1_2, l2_1_2, l3_3, special)
		SELECT 'Иванов Иван Иванович', 'средняя', 85.5, 90.0, 78.3, 85.5, 90.0, 100.0, 95.5, 88.0, 92.5, 75.0, 80.0, 85.0, 90.0, 70.0, 60.0, 65.0, 70.0, 75.0, 101 WHERE NOT EXISTS (SELECT 1 FROM personal_metrics WHERE fio = 'Иванов Иван Иванович');

		INSERT INTO personal_metrics (fio, dadt, smar, smap, smsd, pct_smar, pct_smap, rvp, tpi, rvpd, tpid, gdp, pss, psa, iis, ipe, l0_4, l1_2, l2_1_2, l3_3, special)
		SELECT 'Петрова Мария Сергеевна', 'сильная', 92.1, 88.4, 85.0, 92.1, 88.4, 98.0, 91.2, 86.5, 90.1, 78.5, 82.3, 88.1, 91.0, 75.5, 68.2, 72.1, 78.4, 80.2, 5 WHERE NOT EXISTS (SELECT 1 FROM personal_metrics WHERE fio = 'Петрова Мария Сергеевна');

		INSERT INTO personal_metrics (fio, dadt, smar, smap, smsd, pct_smar, pct_smap, rvp, tpi, rvpd, tpid, gdp, pss, psa, iis, ipe, l0_4, l1_2, l2_1_2, l3_3, special)
		SELECT 'Сидоров Алексей Владимирович', 'слабая', 76.5, 81.2, 70.4, 76.5, 81.2, 95.0, 88.1, 75.4, 85.2, 70.1, 75.4, 80.2, 85.5, 60.1, 55.4, 62.3, 68.1, 71.2, 2 WHERE NOT EXISTS (SELECT 1 FROM personal_metrics WHERE fio = 'Сидоров Алексей Владимирович');

		INSERT INTO personal_metrics (fio, dadt, smar, smap, smsd, pct_smar, pct_smap, rvp, tpi, rvpd, tpid, gdp, pss, psa, iis, ipe, l0_4, l1_2, l2_1_2, l3_3, special)
		SELECT 'Кузнецова Елена Дмитриевна', 'средняя', 88.8, 91.0, 82.5, 88.8, 91.0, 102.0, 96.5, 90.2, 94.1, 80.5, 85.6, 89.4, 93.2, 78.5, 70.1, 75.4, 80.5, 85.1, 10 WHERE NOT EXISTS (SELECT 1 FROM personal_metrics WHERE fio = 'Кузнецова Елена Дмитриевна');

		INSERT INTO personal_metrics (fio, dadt, smar, smap, smsd, pct_smar, pct_smap, rvp, tpi, rvpd, tpid, gdp, pss, psa, iis, ipe, l0_4, l1_2, l2_1_2, l3_3, special)
		SELECT 'Смирнов Дмитрий Андреевич', 'критическая', 65.2, 70.4, 60.1, 65.2, 70.4, 85.0, 78.2, 65.5, 75.4, 65.2, 68.5, 72.1, 78.5, 50.2, 45.1, 52.4, 58.2, 61.5, 0 WHERE NOT EXISTS (SELECT 1 FROM personal_metrics WHERE fio = 'Смирнов Дмитрий Андреевич');

		INSERT INTO personal_metrics (fio, dadt, smar, smap, smsd, pct_smar, pct_smap, rvp, tpi, rvpd, tpid, gdp, pss, psa, iis, ipe, l0_4, l1_2, l2_1_2, l3_3, special)
		SELECT 'Попова Ольга Игоревна', 'сильная', 94.5, 96.2, 90.5, 94.5, 96.2, 105.0, 99.1, 95.4, 98.2, 85.4, 90.2, 94.5, 98.1, 82.5, 78.2, 82.5, 88.4, 92.1, 15 WHERE NOT EXISTS (SELECT 1 FROM personal_metrics WHERE fio = 'Попова Ольга Игоревна');

		INSERT INTO personal_metrics (fio, dadt, smar, smap, smsd, pct_smar, pct_smap, rvp, tpi, rvpd, tpid, gdp, pss, psa, iis, ipe, l0_4, l1_2, l2_1_2, l3_3, special)
		SELECT 'Васильев Сергей Николаевич', 'слабая', 78.5, 82.4, 74.2, 78.5, 82.4, 92.0, 86.5, 78.2, 88.4, 72.5, 78.1, 82.4, 86.5, 65.4, 58.2, 65.1, 71.4, 75.2, 3 WHERE NOT EXISTS (SELECT 1 FROM personal_metrics WHERE fio = 'Васильев Сергей Николаевич');

		INSERT INTO personal_metrics (fio, dadt, smar, smap, smsd, pct_smar, pct_smap, rvp, tpi, rvpd, tpid, gdp, pss, psa, iis, ipe, l0_4, l1_2, l2_1_2, l3_3, special)
		SELECT 'Соколова Анна Павловна', 'средняя', 83.2, 85.6, 78.9, 83.2, 85.6, 96.0, 89.4, 82.5, 90.1, 75.2, 80.5, 84.2, 88.5, 68.9, 62.4, 68.2, 74.5, 78.6, 8 WHERE NOT EXISTS (SELECT 1 FROM personal_metrics WHERE fio = 'Соколова Анна Павловна');

		INSERT INTO personal_metrics (fio, dadt, smar, smap, smsd, pct_smar, pct_smap, rvp, tpi, rvpd, tpid, gdp, pss, psa, iis, ipe, l0_4, l1_2, l2_1_2, l3_3, special)
		SELECT 'Михайлов Павел Олегович', 'слабая', 71.4, 76.8, 68.5, 71.4, 76.8, 88.0, 82.1, 72.4, 80.5, 68.2, 72.5, 76.8, 82.1, 58.5, 52.1, 58.4, 64.2, 68.5, 4 WHERE NOT EXISTS (SELECT 1 FROM personal_metrics WHERE fio = 'Михайлов Павел Олегович');

		INSERT INTO personal_metrics (fio, dadt, smar, smap, smsd, pct_smar, pct_smap, rvp, tpi, rvpd, tpid, gdp, pss, psa, iis, ipe, l0_4, l1_2, l2_1_2, l3_3, special)
		SELECT 'Новикова Татьяна Викторовна', 'сильная', 89.6, 93.1, 85.4, 89.6, 93.1, 101.0, 95.2, 89.5, 95.6, 82.1, 86.4, 90.1, 94.2, 75.6, 68.4, 74.2, 80.1, 84.5, 12 WHERE NOT EXISTS (SELECT 1 FROM personal_metrics WHERE fio = 'Новикова Татьяна Викторовна');

		INSERT INTO personal_metrics (fio, dadt, smar, smap, smsd, pct_smar, pct_smap, rvp, tpi, rvpd, tpid, gdp, pss, psa, iis, ipe, l0_4, l1_2, l2_1_2, l3_3, special)
		SELECT 'Федоров Артем Алексеевич', 'критическая', 68.5, 72.1, 64.2, 68.5, 72.1, 86.0, 79.5, 68.4, 78.2, 66.5, 70.2, 74.5, 79.8, 54.2, 48.5, 54.1, 60.2, 64.5, 1 WHERE NOT EXISTS (SELECT 1 FROM personal_metrics WHERE fio = 'Федоров Артем Алексеевич');

		INSERT INTO personal_metrics (fio, dadt, smar, smap, smsd, pct_smar, pct_smap, rvp, tpi, rvpd, tpid, gdp, pss, psa, iis, ipe, l0_4, l1_2, l2_1_2, l3_3, special)
		SELECT 'Морозова Виктория Александровна', 'сильная', 91.2, 94.5, 88.2, 91.2, 94.5, 103.0, 97.1, 92.4, 96.2, 84.5, 88.2, 91.5, 95.4, 78.2, 72.4, 76.5, 83.1, 86.4, 6 WHERE NOT EXISTS (SELECT 1 FROM personal_metrics WHERE fio = 'Морозова Виктория Александровна');

		INSERT INTO personal_metrics (fio, dadt, smar, smap, smsd, pct_smar, pct_smap, rvp, tpi, rvpd, tpid, gdp, pss, psa, iis, ipe, l0_4, l1_2, l2_1_2, l3_3, special)
		SELECT 'Волков Кирилл Максимович', 'слабая', 74.5, 78.2, 70.1, 74.5, 78.2, 90.0, 84.1, 75.2, 83.4, 70.5, 74.2, 78.5, 83.2, 62.1, 55.4, 62.5, 68.4, 72.1, 7 WHERE NOT EXISTS (SELECT 1 FROM personal_metrics WHERE fio = 'Волков Кирилл Максимович');

		INSERT INTO personal_metrics (fio, dadt, smar, smap, smsd, pct_smar, pct_smap, rvp, tpi, rvpd, tpid, gdp, pss, psa, iis, ipe, l0_4, l1_2, l2_1_2, l3_3, special)
		SELECT 'Павлова Дарья Романовна', 'средняя', 86.4, 89.2, 81.5, 86.4, 89.2, 99.0, 92.4, 86.1, 92.5, 78.2, 82.5, 86.4, 90.5, 72.4, 65.2, 71.4, 77.5, 81.2, 9 WHERE NOT EXISTS (SELECT 1 FROM personal_metrics WHERE fio = 'Павлова Дарья Романовна');

		INSERT INTO personal_metrics (fio, dadt, smar, smap, smsd, pct_smar, pct_smap, rvp, tpi, rvpd, tpid, gdp, pss, psa, iis, ipe, l0_4, l1_2, l2_1_2, l3_3, special)
		SELECT 'Семенов Максим Денисович', 'средняя', 79.8, 83.5, 75.6, 79.8, 83.5, 93.0, 87.2, 80.1, 88.5, 73.8, 78.5, 82.1, 86.4, 66.5, 59.8, 66.4, 72.1, 76.8, 11 WHERE NOT EXISTS (SELECT 1 FROM personal_metrics WHERE fio = 'Семенов Максим Денисович');`

	if _, err := db.Exec(testRows); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("DB created2")
	testRows = `
		INSERT INTO group_reports (institute, group_name, starosta, students_list, students_fact, missed_dln, quality_pct, measured_at, starosta_was)
		SELECT 'ИТ', 'ПИ-21', 'Смирнов Алексей Владимирович', 28, 26, 2, 88.5, '2025-09-10', 'Да' WHERE NOT EXISTS (SELECT 1 FROM group_reports WHERE group_name = 'ПИ-21' AND measured_at = '2025-09-10');

		INSERT INTO group_reports (institute, group_name, starosta, students_list, students_fact, missed_dln, quality_pct, measured_at, starosta_was)
		SELECT 'ИЭ', 'ЭК-22', 'Кузнецова Анна Сергеевна', 25, 24, 1, 95.0, '2025-09-12', 'Нет' WHERE NOT EXISTS (SELECT 1 FROM group_reports WHERE group_name = 'ЭК-22' AND measured_at = '2025-09-12');

		INSERT INTO group_reports (institute, group_name, starosta, students_list, students_fact, missed_dln, quality_pct, measured_at, starosta_was)
		SELECT 'ФЭМ', 'БИ-23', 'Попов Дмитрий Андреевич', 22, 20, 4, 80.5, '2025-09-15', 'Был после' WHERE NOT EXISTS (SELECT 1 FROM group_reports WHERE group_name = 'БИ-23' AND measured_at = '2025-09-15');

		INSERT INTO group_reports (institute, group_name, starosta, students_list, students_fact, missed_dln, quality_pct, measured_at, starosta_was)
		SELECT 'ИЦС', 'ЦПИ-22', 'Соколова Мария Игоревна', 30, 28, 3, 91.2, '2025-10-01', 'Да' WHERE NOT EXISTS (SELECT 1 FROM group_reports WHERE group_name = 'ЦПИ-22' AND measured_at = '2025-10-01');

		INSERT INTO group_reports (institute, group_name, starosta, students_list, students_fact, missed_dln, quality_pct, measured_at, starosta_was)
		SELECT 'ИГ', 'ГЕО-21', 'Лебедев Никита Олегович', 20, 18, 6, 75.0, '2025-10-05', 'Уважительная причина' WHERE NOT EXISTS (SELECT 1 FROM group_reports WHERE group_name = 'ГЕО-21' AND measured_at = '2025-10-05');

		INSERT INTO group_reports (institute, group_name, starosta, students_list, students_fact, missed_dln, quality_pct, measured_at, starosta_was)
		SELECT 'ИТ', 'ИС-22', 'Новикова Елена Павловна', 27, 27, 0, 98.4, '2025-10-12', 'Да' WHERE NOT EXISTS (SELECT 1 FROM group_reports WHERE group_name = 'ИС-22' AND measured_at = '2025-10-12');

		INSERT INTO group_reports (institute, group_name, starosta, students_list, students_fact, missed_dln, quality_pct, measured_at, starosta_was)
		SELECT 'ИЭ', 'МЕН-21', 'Морозов Артем Игоревич', 24, 21, 5, 82.1, '2025-10-20', 'Отсутствовал' WHERE NOT EXISTS (SELECT 1 FROM group_reports WHERE group_name = 'МЕН-21' AND measured_at = '2025-10-20');

		INSERT INTO group_reports (institute, group_name, starosta, students_list, students_fact, missed_dln, quality_pct, measured_at, starosta_was)
		SELECT 'ФЭМ', 'ФИ-22', 'Волкова Ольга Дмитриевна', 26, 25, 1, 94.0, '2025-11-02', 'Да' WHERE NOT EXISTS (SELECT 1 FROM group_reports WHERE group_name = 'ФИ-22' AND measured_at = '2025-11-02');

		INSERT INTO group_reports (institute, group_name, starosta, students_list, students_fact, missed_dln, quality_pct, measured_at, starosta_was)
		SELECT 'ИЦС', 'ИБ-23', 'Зайцев Павел Сергеевич', 23, 20, 7, 78.9, '2025-11-10', 'Был после' WHERE NOT EXISTS (SELECT 1 FROM group_reports WHERE group_name = 'ИБ-23' AND measured_at = '2025-11-10');

		INSERT INTO group_reports (institute, group_name, starosta, students_list, students_fact, missed_dln, quality_pct, measured_at, starosta_was)
		SELECT 'ИГ', 'ЭГ-22', 'Павлова Дарья Романовна', 21, 21, 0, 99.1, '2025-11-15', 'Да' WHERE NOT EXISTS (SELECT 1 FROM group_reports WHERE group_name = 'ЭГ-22' AND measured_at = '2025-11-15');
	`
	if _, err := db.Exec(testRows); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("DB created3")
	fmt.Println(err)
	return nil
}
