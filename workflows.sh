# read the workflow template
WORKFLOW_TEMPLATE=$(cat .github/workflow-template.yml)

mkdir -p .github/workflows

# iterate each go package in the src directory
for package in $(ls src); do
	echo "generating workflow for src/${package}"

	# replace template package placeholder with package name
	WORKFLOW=$(echo "${WORKFLOW_TEMPLATE}" | sed "s/{{package}}/${package}/g")
	
	# save workflow to .github/workflows/{package}
	echo "${WORKFLOW}" > .github/workflows/${package}.yml
done