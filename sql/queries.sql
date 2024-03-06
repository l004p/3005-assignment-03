-- name: GetAllStudents :many
SELECT *
FROM students;

-- name: AddStudent :one
INSERT INTO students (
    first_name, last_name, email, enrollment_date
    ) VALUES
    ($1, $2, $3, $4)
    RETURNING *;

-- name: UpdateStudent :exec
UPDATE students
SET email = $2
WHERE student_id = $1;

-- name: DeleteStudent :exec
DELETE FROM students
WHERE student_id = $1;