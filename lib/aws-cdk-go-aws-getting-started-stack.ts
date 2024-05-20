import * as cdk from "aws-cdk-lib";
import { LambdaIntegration, RestApi } from "aws-cdk-lib/aws-apigateway";
import { Code, Function, Runtime } from "aws-cdk-lib/aws-lambda";
import { Construct } from "constructs";

export class AwsCdkGoAwsGettingStartedStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const myFunc = new Function(this, "HelloHandler", {
      code: Code.fromAsset("lambdas"),
      handler: "main",
      runtime: Runtime.PROVIDED_AL2023,
    });

    const gateway = new RestApi(this, "HelloGateway", {
      defaultCorsPreflightOptions: {
        allowOrigins: ["*"],
        allowMethods: ["GET", "POST", "PUT", "DELETE", "OPTIONS"],
        allowHeaders: ["*"],
      },
    });

    const integration = new LambdaIntegration(myFunc);
    const testResource = gateway.root.addResource("test");
    testResource.addMethod("GET", integration);
  }
}
