UPDATE tasks
SET name=$1, description=$2, deadline_date=$3, completion_status=$4
WHERE task_id=$5
