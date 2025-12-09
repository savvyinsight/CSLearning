cmake_minimum_required(VERSION 3.10)

project(tcp_chat)

set(CMAKE_C_STANDARD 11)
set(CMAKE_C_STANDARD_REQUIRED ON)

# Build both executables
add_executable(client tcp_client.c)
add_executable(serverWThread tcp_serverwith_thread.c)
add_executable(server tcp_server.c)
add_executable(server_udp udp_server.c)
add_executable(client_udp udp_client.c)

# Link libraries (both need socket)
target_link_libraries(client)
target_link_libraries(server)