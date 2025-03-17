在 Kubernetes 中，StatefulSet 是一种用于管理有状态应用的控制器，它的更新规则与 Deployment 等无状态控制器有所不同。StatefulSet 的更新规则主要围绕其设计目标：确保 Pod 的顺序性、唯一性和状态一致性。以下是关于 StatefulSet 更新规则的详细说明：

### 1. **更新策略 (Update Strategy)**

StatefulSet 支持两种主要的更新策略，通过 `.spec.updateStrategy` 字段配置：

- **OnDelete**: 默认策略。在这种模式下，StatefulSet 不会自动更新 Pod。只有当 Pod 被手动删除后，StatefulSet 才会根据最新的模板重新创建 Pod。这意味着更新需要手动干预。
- **RollingUpdate**: 滚动更新策略。StatefulSet 会按照 Pod 的顺序（从最高索引到最低索引，例如 `pod-n` 到 `pod-0`）逐步更新 Pod。每次更新一个 Pod，确保前一个 Pod 完成更新并进入 `Ready` 状态后，再更新下一个。

可以通过以下 YAML 示例指定更新策略：

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: example
spec:
  replicas: 3
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      partition: 0  # 可选，见下文
  template:
    spec:
      containers:
      - name: app
        image: my-app:latest
```

### 2. **RollingUpdate 的细节**

- **顺序更新**: StatefulSet 按 Pod 的序号从高到低更新（例如 `pod-2` -> `pod-1` -> `pod-0`），以确保依赖关系的正确性。
- **逐步进行**: 每次只更新一个 Pod，只有当当前 Pod 更新完成并变为 `Ready` 状态后，才会继续更新下一个。
- **Partition（分区）**: 可选字段 `.spec.updateStrategy.rollingUpdate.partition` 用于控制更新的范围。如果设置了 `partition`，则只有序号大于或等于 `partition` 的 Pod 会被更新，小于 `partition` 的 Pod 保持不变。例如：
  - 如果 `partition: 1`，且有 3 个 Pod（`pod-0`, `pod-1`, `pod-2`），则只更新 `pod-1` 和 `pod-2`，`pod-0` 不变。
  - 这在分阶段部署或金丝雀发布时非常有用。

### 3. **注意事项**

- **不可随意更改 Pod 名称或数量**: StatefulSet 的 Pod 具有固定名称（例如 `name-0`, `name-1`），更新时不能随意更改 `.spec.replicas` 或 Pod 模板中的关键字段（如 `volumeClaimTemplates`），否则可能会导致错误。
- **PVC（持久卷声明）不变**: StatefulSet 的每个 Pod 通常绑定一个持久卷声明 (PVC)，更新时不会重新生成 PVC，因此存储状态保持一致。
- **手动干预**: 如果更新失败（例如新镜像不可用），StatefulSet 不会自动回滚，需要手动修复或删除有问题的 Pod。

### 4. **示例：触发滚动更新**

假设你修改了 StatefulSet 的镜像：

```yaml
spec:
  template:
    spec:
      containers:
      - name: app
        image: my-app:v2  # 从 v1 更新到 v2
```

应用后，Kubernetes 会根据 `RollingUpdate` 策略依次更新 Pod，确保服务不中断。

### 总结

StatefulSet 的更新规则旨在保证有状态应用的稳定性和顺序性。默认情况下使用 `OnDelete` 策略，需要手动删除 Pod 触发更新；而 `RollingUpdate` 提供自动化滚动更新，支持分区控制。
