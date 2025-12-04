cmake_minimum_required(version 3.10)

project(tcp_chat)

set(CMAKE_C_STANDARD 11)
set(CMAKE_C_STANDARD_REQUIRED ON)

# Build both executables
add_executable(client tcp_client.c)
add_executable(server tcp_server.c)

# Link libraries (both need socket)
target_link_libraries(client)
target_link_libraries(server)