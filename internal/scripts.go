package internal

var (
	LoadToVps = `
      img="$1"
      user="$2"
      host="$3"
      key="$4"
      
      # debugging
      echo "Image: $img"
      echo "User: $user"
      echo "Host: $host"
      

      docker save -o "${img}.tar" "$img"
      scp -i "${key}" -o StrictHostKeyChecking=no "${img}.tar" "${user}@${host}:/home/${user}/"
  `
)
