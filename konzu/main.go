package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "time"

    pb "konzu/grpc" // Paquete generado por protoc
    //"clients/utils" // Paquete para utilidades (utils.go)
    "google.golang.org/grpc"
)

// Estructura para almacenar el estado de los paquetes
type PackageInfo struct {
    OrderID   string
    Status    string
    Attempts  int
    Value     float64
    Client    string
    Item      string
    Destination string
    Caravana  string
}

var packageRegistry = make(map[string]*PackageInfo)

type server struct {
    pb.UnimplementedLogisticsServiceServer
}

// Función para procesar la orden
func (s *server) ProcessOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
    log.Printf("Recibida orden del cliente %s para el item %s con cantidad %d", req.ClientName, req.ItemName, req.Quantity)
    
    // Simulamos la creación de un ID de orden único
    orderID := fmt.Sprintf("ORD-%d", time.Now().UnixNano())
    
    // Crear un registro para la orden
    packageInfo := &PackageInfo{
        OrderID:   orderID,
        Status:    "En Cetus",
        Attempts:  0,
        Value:     float64(req.Value),
        Client:    req.ClientName,
        Item:      req.ItemName,
        Destination: req.Destination,
        Caravana:  "Pendiente",
    }
    packageRegistry[orderID] = packageInfo

    log.Printf("Orden registrada con ID %s para destino %s", orderID, req.Destination)
    return &pb.OrderResponse{
        OrderId: orderID,
        Message: "Orden procesada exitosamente",
    }, nil
}

// Función para consultar el estado de una orden
func (s *server) CheckStatus(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {
    log.Printf("Consultando estado de la orden %s", req.OrderId)
    
    // Verificar si la orden existe
    packageInfo, exists := packageRegistry[req.OrderId]
    if !exists {
        return &pb.StatusResponse{
            OrderId: req.OrderId,
            Status:  "No Encontrado",
            Message: "No se encontró la orden con ese código",
        }, nil
    }

    return &pb.StatusResponse{
        OrderId: req.OrderId,
        Status:  packageInfo.Status,
        Message: "Estado del paquete obtenido correctamente",
    }, nil
}

func main() {
    // Iniciar el servidor gRPC
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterLogisticsServiceServer(grpcServer, &server{})

    log.Println("Servidor de Logística/Konzu iniciado en el puerto 50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
    }
}
