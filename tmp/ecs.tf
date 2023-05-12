resource "aws_ecs_cluster" "local_stack" {
  name = "local-stack"
}

resource "aws_ecs_task_definition" "local_stack" {
  family                   = "local-stack"
  container_definitions    = file("./container_definitions/local_stack.json")
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  task_role_arn            = aws_iam_role.ecs_task_role.arn
}

resource "aws_ecs_service" "local_stack" {
  name            = "local-stack"
  cluster         = aws_ecs_cluster.local_stack.id
  task_definition = aws_ecs_task_definition.local_stack.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = [module.vpc.private_subnets[0]]
    security_groups  = [aws_security_group.local_stack.id]
    assign_public_ip = true
  }
}

resource "aws_security_group" "local_stack" {
  name        = "local-stack"
  description = "local-stack"
  vpc_id      = module.vpc.vpc_id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
